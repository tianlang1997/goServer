package redisClient

import (
	"github.com/chasex/redis-go-cluster"
	"time"
	"../../config"
)

var messageClient,_ = redis.NewCluster(&redis.Options{
	StartNodes:   config.RedisClient.GetMessage(),
	ConnTimeout:  50 * time.Millisecond,
	ReadTimeout:  50 * time.Millisecond,
	WriteTimeout: 50 * time.Millisecond,
	KeepAlive:    16,
	AliveTime:    60 * time.Second,
})
var commonClient,_ = redis.NewCluster(&redis.Options{
	StartNodes:   config.RedisClient.GetCommon(),
	ConnTimeout:  50 * time.Millisecond,
	ReadTimeout:  50 * time.Millisecond,
	WriteTimeout: 50 * time.Millisecond,
	KeepAlive:    16,
	AliveTime:    60 * time.Second,
})

func GetMessageClien()(*redis.Cluster)  {
	return  messageClient
}

func GetCommonClien()(*redis.Cluster)  {
	return  commonClient
}
