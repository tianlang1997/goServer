package cache

import (
	"goServer.com/m/connector/redisClient"
)

var cache = NewStringCache("test", redisClient.GetCommonClient(), 60)

type testCache struct {
}

func (this *testCache) GetVal(key string) (interface{}, error) {
	return cache.Get(key)
}

func (this *testCache) SetVal(key, val string) error {
	return cache.Set(key, val)
}
