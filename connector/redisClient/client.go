package redisClient

import (
	"../../config"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	clusterClient *redis.ClusterClient
	client *redis.Client
}

func (this *RedisClient) GetClient(args...interface{})  *redis.Client {
	return this.client
}

func (this *RedisClient) GetClusterClient(args...interface{})  *redis.ClusterClient {
	return this.clusterClient
}

func (this *RedisClient) Do(args...interface{})  *redis.Cmd {
	if this.clusterClient != nil {
		return this.clusterClient.Do(args)
	}
	if this.client != nil {
		return this.client.Do(args)
	}
	panic("redis no ready!")
	return nil
}

func (this *RedisClient) Get(key string) *redis.StringCmd {
	if this.clusterClient != nil {
		return this.clusterClient.Get(key)
	}
	if this.client != nil {
		return this.client.Get(key)
	}
	panic("redis no ready!")
	return nil
}

func (this *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if this.clusterClient != nil {
		return this.clusterClient.Set(key,value,expiration)
	}
	if this.client != nil {
		return this.client.Set(key,value,expiration)
	}
	panic("redis no ready!")
	return nil
}

func (this *RedisClient) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if this.clusterClient != nil {
		return this.clusterClient.SetNX(key,value,expiration)
	}
	if this.client != nil {
		return this.client.SetNX(key,value,expiration)
	}
	panic("redis no ready!")
	return nil
}



var messageClient,commonClient *RedisClient
func GetMessageClient() *RedisClient {
	return messageClient
}

func GetCommonClient() *RedisClient  {
	return commonClient
}


func init()  {
	messageClient = &RedisClient{}
	initClient(messageClient,config.RedisClient.GetMessage())

	commonClient = &RedisClient{}
	initClient(commonClient,config.RedisClient.GetMessage())
}

func initClient(client *RedisClient,addrs []string)  {
	client.clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		DialTimeout: 50 * time.Millisecond,
		PoolSize: 100,
	})

	pong, err := client.clusterClient.Ping().Result()
	fmt.Println(pong, err)
	if err != nil{
		panic(err)

		client.client = redis.NewClient(&redis.Options{
			Addr:     addrs[0],
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong, err := client.client.Ping().Result()
		fmt.Println(pong, err)
		if err != nil {
			panic(err)
		}
	}
}