package redis

import (
	"fmt"
	"gopkg.in/redis.v4"
)

func ConnectRedis(serverAddr []string) *redis.ClusterClient {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: serverAddr,
	})

	_, err := client.Ping().Result()
	if err == nil{
		fmt.Println("redis connect ok")
	}

	return client
}