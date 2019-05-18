package config

type redisClient struct {
}

var RedisClient = redisClient{}

func (rc *redisClient) GetMessage() []string {
	return []string{"0.0.0.0:7000", "0.0.0.0:7001", "0.0.0.0:7002", "0.0.0.0:7003",  "0.0.0.0:7005"}
}

func (rc *redisClient) GetCommon() []string {
	return []string{"0.0.0.0:7000", "0.0.0.0:7001", "0.0.0.0:7003", "0.0.0.0:7004", "0.0.0.0:7005"}
}

func (rc *redisClient) GetSession() []string {
	return []string{"0.0.0.0:7000", "0.0.0.0:7001", "0.0.0.0:7003", "0.0.0.0:7004", "0.0.0.0:7005"}
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
