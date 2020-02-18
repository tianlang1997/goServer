package config

import (
	"fmt"
	"goServer.com/m/module"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

var env="default"
var envMap= map[string]string{"defalut":"goServer.com/m/config/defalut.yml","product":"goServer.com/m/config/product.yml"}
var config = module.Config{}
var BasePath string
func Init(e string)  {
	_, filename, _, _ := runtime.Caller(1)
	BasePath = path.Dir(filename)
	if envMap[e] != ""	{
		env=e
	}
	configPath := path.Join(BasePath,"./config/default.yml")
	if(env == "product" ){
		configPath = path.Join(BasePath,"./config/product.yml")
	}

	data ,rErr := ioutil.ReadFile(configPath)
	if rErr!=nil{
		panic(rErr)
	}
	yErr := yaml.Unmarshal([]byte(data), &config)
	if yErr != nil {
		panic(yErr)
	}
	fmt.Println("session addr:",config.Redis.Session.Addr)
}
type redisClient struct {
}

var RedisClient = redisClient{}

func (rc *redisClient) GetMessage() []string {
	return []string{"0.0.0.0:6379", "0.0.0.0:6379", "0.0.0.0:7002", "0.0.0.0:7003",  "0.0.0.0:7005"}
}

func (rc *redisClient) GetCommon() []string {
	return []string{"0.0.0.0:6379", "0.0.0.0:7001", "0.0.0.0:7003", "0.0.0.0:7004", "0.0.0.0:7005"}
}

func (rc *redisClient) GetSession() []string {
	return []string{"0.0.0.0:6379", "0.0.0.0:7001", "0.0.0.0:7003", "0.0.0.0:7004", "0.0.0.0:7005"}
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
