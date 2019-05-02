package datastruct

type User struct {
	Name string `bson:"name,omitempty"`
	Age  int32  `bson:"age,omitempty"`
}
