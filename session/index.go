package session

import (
	"../config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"time"
)
var Session = sessions.New(sessions.Config{
	Cookie:  "sessionscookieid",
	Expires: 7 * 24 * time.Hour, // 0 代表忽略
})

func init()  {
	db := redis.New(service.Config{
		Addr:config.RedisClient.GetSession()[0],
	})
	// 按下control + C / cmd + C时关闭并解锁数据库
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})
	// 重要:
	Session.UseDatabase(db)

}

