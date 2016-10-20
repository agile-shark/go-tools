package redis

import (
"github.com/chasex/redis-go-cluster"
"time"
"fmt"
)

func ConnectRedis1(serverAddr []string) *redis.Cluster {

  cluster, err := redis.NewCluster(
    &redis.Options{
      StartNodes: serverAddr,
      ConnTimeout: 50 * time.Millisecond,
      ReadTimeout: 50 * time.Millisecond,
      WriteTimeout: 50 * time.Millisecond,
      KeepAlive: 16,
      AliveTime: 60 * time.Second,
    })
  if err == nil{
    fmt.Println("redis connect ok")
  }
  return cluster
}
