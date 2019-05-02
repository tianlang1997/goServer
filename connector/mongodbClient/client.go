package mongodbClient

import (
	"log"
)
import (
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session

func init() {
	//var configInfo = config.MongodbClient.GetMongoDBConfig();
	//
	//dialInfo := &mgo.DialInfo{
	//	Addrs:     []string{configInfo.DBhost}, //数据库地址 dbhost: mongodb://user@123456:127.0.0.1:27017
	//	Timeout:   60 * time.Second,                     // 连接超时时间 timeout: 60 * time.Second
	//	Source:    configInfo.Authdb,                     // 设置权限的数据库 authdb: admin
	//	Username:  configInfo.Authuser,                 // 设置的用户名 authuser: user
	//	Password:  configInfo.Authpass,                // 设置的密码 authpass: 123456
	//	PoolLimit: configInfo.Poollimit,       // 连接池的数量 poollimit: 100
	//}
	//
	//s, err := mgo.DialWithInfo(dialInfo)
	s, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(doc)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

//更新，如果不存在就插入一个新的数据 `upsert:true`
func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

// `multi:true`
func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}
func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}
