package config

type redisClient struct {
}

var RedisClient = redisClient{}

func (rc *redisClient) GetMessage() []string {
	return []string{"192.168.159.137:7000", "192.168.159.137:7001", "192.168.159.137:7002", "192.168.159.137:7003", "192.168.159.137:7004", "192.168.159.137:7005"}
}

func (rc *redisClient) GetCommon() []string {
	return []string{"192.168.159.137:7000", "192.168.159.137:7001", "192.168.159.137:7002", "192.168.159.137:7003", "192.168.159.137:7004", "192.168.159.137:7005"}
}

type mongodbClient struct {
	DBhost    string
	Timeout   int
	Authdb    string
	Authuser  string
	Authpass  string
	Poollimit int
}
type MongodbClientConfig struct {
	DBhost    string
	Authdb    string
	Authuser  string
	Authpass  string
	Poollimit int
}

var MongodbClient = mongodbClient{}

func (this *mongodbClient) GetMongoDBConfig() *MongodbClientConfig {
	return &MongodbClientConfig{
		DBhost:    "localhost:27017",
		Authdb:    "admin",
		Authuser:  "admin",
		Authpass:  "",
		Poollimit: 100,
	}
}
