package nsq

import (
	"testing"
)

func Test_consumers(t *testing.T){
	Consumer()
}

func Test_producer(t *testing.T){
	msgQueue := &MsgQueue{}
	msgQueue.Init("173.199.124.134:4150")
	for{
		msgQueue.Producer("test1", []string{"aaaa"})
	}
}