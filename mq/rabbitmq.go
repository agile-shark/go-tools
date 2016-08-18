// rabbitmq project rabbitmq.go
package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

const (
	USER_NAME     = "guest"
	USER_PASSWORD = "guest"
)

type MqQueue struct {
	userName     string
	userPassword string
	serverAddr   string
	queueName    string
	exchange     string
	routeKey     string
}

func NewMqQueue(serverAddr, queueName, exchange, routeKey string) *MqQueue {
	return &MqQueue{userName: USER_NAME, userPassword: USER_PASSWORD, serverAddr: serverAddr,
		queueName: queueName, exchange: exchange, routeKey: routeKey}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func GetStrReadMQ(mqQueue *MqQueue, limit int, autoAck bool) (result []string){

	result = make([]string, limit)
	messages := consumer(mqQueue, limit)
	for i, delivery := range messages {
		//log.Printf("Received a message : %s, count : %d", delivery.Body, i)
		log.Printf("Received count : %d", i)
		result = append(result, string(delivery.Body))
		//err := delivery.Ack(autoAck)
		//failOnError(err, "ACK RabbitMQ error")
	}
	return result
}

func consumer(mqQueue *MqQueue, limit int) (result []amqp.Delivery){
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
		mqQueue.queueName,	// name
		false, 				// durable
		false, 				// delete when usused
		false,  				// exclusive
		false, 				// no-wait
		nil,   				// arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		mqQueue.queueName, 		// queue name
		mqQueue.routeKey,     	// routing key
		mqQueue.exchange,     	// exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue") // delete when usused

	msgs, err := ch.Consume(
		q.Name, 			// queue
		"",     			// consumer
		true,   			// auto-ack
		false,  			// exclusive
		false,  			// no-local
		false,  			// no-wait
		nil,    			// args
	)
	failOnError(err, "Failed to register a consumer")

	//forever := make(chan bool)
	result = []amqp.Delivery{}
	i := 1
	for d := range msgs {
//		log.Printf("Received a message : %s, count : %d", d.Body, i)
		result = append(result, d)
		if i < limit {
			i++
		}else {
			break
		}
	}
	return result
	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}

func Counts(mqQueue *MqQueue) (counts int) {

	//eg."amqp://guest:guest@10.1.4.83:5672/"
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

	q, err := ch.QueueDeclare(
		mqQueue.queueName,    	// name
		false, 					// durable
		false, 					// delete when usused
		false,  				// exclusive
		false, 					// no-wait
		nil,   					// arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("queueName : %s, count : %d", mqQueue.queueName, q.Messages)

	ch.Close()

	return q.Messages
}

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

	_, err = ch.QueueDeclare(
		mqQueue.queueName,    	// name
		false, 					// durable
		false, 					// delete when usused
		false,  				// exclusive
		false, 					// no-wait
		nil,   					// arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		mqQueue.queueName,    // queue name
		mqQueue.routeKey,     // routing key
		mqQueue.exchange,     // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue") // delete when usused

//	body := bodyFrom(message)
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

		log.Printf(" [x] Sent %s", message)
	}
	ch.Close()
}

//func bodyFrom(args []string) string {
//	return strings.Join(args, " ")
//}
