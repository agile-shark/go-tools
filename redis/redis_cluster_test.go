package redis

import (
	"testing"
)

func TestConnectRedis(t *testing.T){

	for{
		client := ConnectRedis([]string{"10.1.4.106:6401", "10.1.4.104:6386"})
		client.FlushDb()
		client.Close()

		client1 := ConnectRedis([]string{"10.1.4.104:6389", "10.1.4.105:6408", "10.1.4.105:6409", "10.1.4.106:6402", "10.1.4.106:6403", "10.1.4.104:6390"})
		client1.FlushDb()
		client1.Close()

		client2 := ConnectRedis([]string{"10.1.4.91:6401", "10.1.4.93:6394", "10.1.4.93:6395", "10.1.4.94:6398", "10.1.4.94:6399", "10.1.4.91:6402"})
		client2.FlushDb()
		client2.Close()

		client3 := ConnectRedis([]string{"10.1.4.91:6403", "10.1.4.93:6396", "10.1.4.93:6397", "10.1.4.94:6400", "10.1.4.94:6401", "10.1.4.91:6404"})
		client3.FlushDb()
		client3.Close()
	}

}


//func TestRedis(t *testing.T){
//
//	client := ConnectRedis([]string{"10.1.4.104:6379", "10.1.4.105:6400", "10.1.4.105:6401", "10.1.4.106:6400", "10.1.4.106:6401", "10.1.4.104:6386"})
//
//  	err := client.Set("key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := client.Get("key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := client.Get("key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exists")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}  // Output: key value  // key2 does not exists
//}
