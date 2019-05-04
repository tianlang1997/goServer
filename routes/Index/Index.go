package Index

import (
	"../../cache"
	_ "../../config"
	"../../session"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		//id, _ := ctx.Params().GetInt("id")
		//var err = database.GetUserData().Insert(&datastruct.User{Name:"Bob",Age:19})
		//if err != nil {
		//	panic(err)
		//	str := fmt.Sprintf("%s", err)
		//	ctx.Write([]byte(str))
		//	return
		//}
		//err, ret := database.GetUserData().GetByName("Jack")
		//if err != nil {
		//	fmt.Println("query err ",err)
		//}else {
		//	fmt.Println("query success",ret.Name,ret.Age)
		//}

		_ = ctx.View("Index.html", map[string]string{"word": "test"})
	})
	app.Get("/testJson", func(context iris.Context) {
		//str := []byte(`{"name": "test"}`)
		//context.Write(str)

		//context.JSON(bson.M{"name": "test"})
		mjson, _ := json.Marshal(map[string]string{"name": "test"})
		context.Write(mjson)
	})

	app.Get("/testRedis", func(context iris.Context) {
		err0 := cache.TestCache.SetVal("aaa", "bbb")
		fmt.Println(err0)
		str, err := cache.TestCache.GetVal("aaa")
		fmt.Println(str, err)
		mjson, _ := json.Marshal(map[string]string{"name": str.(string)})
		context.Write(mjson)
	})

	app.Get("/testSetSession", func(context iris.Context) {
		s := session.Session.Start(context)
		//设置
		s.Set("name", "iris")
		context.Write([]byte("ok"))
	})

	app.Get("/testGetSession", func(context iris.Context) {
		s := session.Session.Start(context)
		//设置
		str:=s.Get("name")
		context.Write([]byte(str.(string)))
	})
}
