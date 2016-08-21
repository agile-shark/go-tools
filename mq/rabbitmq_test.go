// rabbitmq project rabbitmq.go
package mq

import(
    "testing"
    "log"
)

func controller(content string) (bool, error) {
    log.Printf("controller Received a message : %s", content)
    return true,nil
}


func Test_consumers(t *testing.T){
    for{
        Counts(NewMqQueue("10.38.14.105:5672", "exchange", "routekey", []string{"queueTest"}))
        Consumer(NewMqQueue("10.38.14.105:5672", "exchange", "routekey", []string{"queueTest1", "queueTest1"}), controller)
        Counts(NewMqQueue("10.38.14.105:5672", "exchange", "routekey", []string{"queueTest"}))
    }
}

func Test_producer(t *testing.T){
    for{
        Producer(NewMqQueue("10.38.14.105:5672", "exchange", "routekey", []string{"queueTest"}), []string{"hello", "world"})
    }
}

func Test_counts(t *testing.T){
    for{
        Counts(NewMqQueue("10.38.14.105:5672", "exchange", "routekey", []string{"queueTest"}))
    }
}