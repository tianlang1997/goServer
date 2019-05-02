package cache

import (
	"../connector/redisClient"
	"time"
)

type StringCache struct {
	base
}

func (this *StringCache) Get(key string) (interface{}, error) {
	return this.client.Get(this.GetKey(key)).Result()
}

func (this *StringCache) Set(key, val string) error {
	err := this.client.Set(this.GetKey(key), val, time.Duration(this.expires)*time.Second).Err()
	return err
}

func NewStringCache(prefix string, client *redisClient.RedisClient, expires int) *StringCache {
	return &StringCache{
		base{
			prefix:  prefix,
			client:  client,
			expires: expires,
		},
	}
}
