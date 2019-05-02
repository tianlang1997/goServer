package config

type redisClient struct {

}
var RedisClient = redisClient{}
func (rc *redisClient)GetMessage()([]string) {
	return []string{"127.0.0.1:6379","127.0.0.1:6379","127.0.0.1:6379"}
}

func (rc *redisClient)GetCommon()([]string) {
	return []string{"127.0.0.1:6379","127.0.0.1:6379","127.0.0.1:6379"}
}