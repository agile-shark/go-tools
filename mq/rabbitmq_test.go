// rabbitmq project rabbitmq.go
package mq

import(
    "testing"
)

func Test_consumers(t *testing.T){
    //for{
   Counts(NewMqQueue("10.38.10.57:5672", "PRE_ARTICLE", "PRE_ARTICLE_EXCHANGE", "pre_article"))
    Counts(NewMqQueue("10.38.11.220:5672", "ARTICLE_UNION_NEW", "ARTICLE_EXCHANGE_NEW", "article.union_new"))
    //GetStrReadMQ(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"), 10)
   // Counts(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"))
   // GetStrReadMQ(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"), 10, true)
    //Counts(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"))
    //}
}

func Test_counts(t *testing.T){
    for{
        Counts(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"))
    }
}

func Test_producer(t *testing.T){
    for{
        Producer(NewMqQueue("10.1.4.83:5672", "queueTest", "exchange", "routekey"), []string{"hello", "world"})
    }
}