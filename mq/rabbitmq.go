// rabbitmq project rabbitmq.go
package mq

import (
	"github.com/streadway/amqp"
	"log"
	"strings"
	"fmt"
)

const (
	USER_NAME     = "guest"
	USER_PASSWORD = "guest"
)

type functionType func(string) (bool, error) // 声明了一个函数类型

type MqQueue struct {
	userName     string
	userPassword string
	serverAddr   string
	QueueName    []string
	exchange     string
	routeKey     string
}

func NewMqQueue(serverAddr, exchange, routeKey string, queueName []string) *MqQueue {
	return &MqQueue{userName: USER_NAME, userPassword: USER_PASSWORD, serverAddr: serverAddr,
		QueueName: queueName, exchange: exchange, routeKey: routeKey}
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Sprintf("%s: %s", msg, err)
		//		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//消费者
func Consumer(mqQueue *MqQueue, controller functionType){

	//eg."amqp://guest:guest@10.1.4.83:5672/"
	url := strings.Join([]string{"amqp://", mqQueue.userName, ":", mqQueue.userPassword, "@", mqQueue.serverAddr}, "")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		mqQueue.exchange,   // name
		"topic", 			// type
		false,     			// durable
		false,    			// auto-deleted
		false,    			// internal
		false,    			// no-wait
		nil,      			// arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		mqQueue.QueueName[0],	// name
		false, 				// durable
		false, 				// delete when usused
		false,  			// exclusive
		false, 				// no-wait
		nil,   				// arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		mqQueue.QueueName[0],   // queue name
		mqQueue.routeKey,     	// routing key
		mqQueue.exchange,     	// exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

//	err = ch.Qos(
//		1,     // prefetch count
//		0,     // prefetch size
//		false, // global
//	)
//	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, 			// queue
		"",     			// consumer
		false,   			// auto-ack
		false,  			// exclusive
		false,  			// no-local
		false,  			// no-wait
		nil,    			// args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	for d := range msgs {

		//log.Printf("Received a message : %s", d.Body)

		ok, err := controller(string(d.Body))
		if ok {
			err1 := d.Ack(true)
			failOnError(err1, "ACK RabbitMQ error")
		}else {
			failOnError(err, "controll data error")
			d.Nack(true, true)
		}
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

//计数
func Counts(mqQueue *MqQueue) (counts map[string]int, err error) {

	//eg."amqp://guest:guest@10.1.4.83:5672/"
	url := strings.Join([]string{"amqp://", mqQueue.userName, ":", mqQueue.userPassword, "@", mqQueue.serverAddr}, "")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	if(err != nil){
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	if(err != nil){
		return nil, err
	}
	defer ch.Close()

	counts = make(map[string]int)
	for _, queueName := range mqQueue.QueueName {

		q, err := ch.QueueDeclare(
			queueName, // name
			false, // durable
			false, // delete when usused
			false, // exclusive
			false, // no-wait
			nil, // arguments
		)
		failOnError(err, "Failed to declare a queue")
		if(err != nil){
			return nil, err
		}
		log.Printf("queueName : %s, count : %d", queueName, q.Messages)

		counts[queueName] = q.Messages
	}
	return counts, nil

}

//生产者
func Producer(mqQueue *MqQueue, messages []string) {

	url := strings.Join([]string{"amqp://", mqQueue.userName, ":", mqQueue.userPassword, "@", mqQueue.serverAddr}, "")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		mqQueue.exchange,    // name
		"topic",             // type
		false,               // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	for _, queueName := range mqQueue.QueueName {

		_, err = ch.QueueDeclare(
			queueName,    	        // name
			false, 					// durable
			false, 					// delete when usused
			false,  				// exclusive
			false, 					// no-wait
			nil,   					// arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = ch.QueueBind(
			queueName,            // queue name
			mqQueue.routeKey,     // routing key
			mqQueue.exchange,     // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue") // delete when usused
	}

	for _, message := range messages  {

		err = ch.Publish(
			mqQueue.exchange,   // exchange
			mqQueue.routeKey,   // routing key
			false,            	// mandatory
			false,            	// immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		failOnError(err, "Failed to publish a message")
	}
}
