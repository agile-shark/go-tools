package nsq

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)


type MsgQueue struct {
	addr string
	producer *nsq.Producer
}

func (this *MsgQueue) Init(addr string) error {
	var err error
	this.addr = addr

	//  try to connect
	cfg := nsq.NewConfig()
	this.producer, err = nsq.NewProducer(addr, cfg)
	if nil != err {
		return err
	}

	//  try to ping
	err = this.producer.Ping()
	if nil != err {
		this.producer.Stop()
		this.producer = nil
		return err
	}
	return nil
}

func (this *MsgQueue) Producer(topic string, messages []string) error {
	this.producer.Publish(topic, []byte(messages[0]))
	return nil
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {
	log.Println("recv:", string(message.Body))
	return nil
}

func Consumer() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test1", "nsq", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQD("173.199.124.134:4150")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}
