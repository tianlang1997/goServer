package cache
import (
	"github.com/chasex/redis-go-cluster"
)
type base struct {
	prefix string
	client *redis.Cluster
	name string
	expires int
}

func (this *base) GetPrefix()(string) {
	return this.prefix
}

func (this *base) GetKey(id string)(string) {
	return this.prefix+id
}

func (this *base) Exec(cmd string,args ...interface{})(interface{},error) {
	return this.client.Do(cmd,args);
}

func (this *base) Expire(key string) (interface{},error) {
	return this.client.Do("expire",key,this.expires)
}

func (this *base) ExpireByTime(key string,expireTime int)(interface{},error)  {
	return this.client.Do("expire",key,expireTime)
}

func (this *base) Release(key string) (interface{},error) {
	return this.client.Do("del",key);
}

func (this *base) Keys(keyPrefix string) (interface{},error) {
	return this.client.Do("keys",keyPrefix);
}