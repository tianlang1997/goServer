package cache

import (
	"../connector/redisClient"
)

type base struct {
	prefix  string
	client  *redisClient.RedisClient
	expires int
}

func (this *base) GetPrefix() string {
	return this.prefix
}

func (this *base) GetKey(id string) string {
	return this.prefix + ":" + id
}
