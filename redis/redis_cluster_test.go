package redis

import (
	"testing"
	"fmt"
)

func TestConnectRedis(t *testing.T){
	cluster := ConnectRedis([]string{"10.1.4.103:6393", "10.1.4.104:6392", "10.1.4.105:6391"})
	for  {
		if cluster != nil {
			fmt.Println("connect ok")
		}
		cluster.Do("SET", "kev", "value")
	}
	cluster.Close()
}
