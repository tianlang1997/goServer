package database

import (
	"goServer.com/m/connector/mongodbClient"
	"goServer.com/m/datastruct"
	"gopkg.in/mgo.v2/bson"
)

type userData struct {
	db string
	collection string
}

func (this * userData) GetByName(name string)(error,datastruct.User)  {
	var result datastruct.User
	var err = mongodbClient.FindOne(this.db,this.collection,bson.M{"name": name}, bson.M{"_id":0}, &result)
	return err,result
}

func (this * userData) Insert(user *datastruct.User) error  {
	var err = mongodbClient.Insert(this.db,this.collection,user);
	return err
}
var userDataInstance *userData

func init()  {
	userDataInstance=&userData{
		db:"testdb",
		collection:"user",
	}
}
func GetUserData()(* userData)  {
	return userDataInstance
}