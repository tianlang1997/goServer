package Index

import (
	_ "../../config"
	"github.com/kataras/iris"
)

func Register(app *iris.Application)  {
	app.Get("/", func(ctx iris.Context) {
		//id, _ := ctx.Params().GetInt("id")
		_ = ctx.View("Index.html")
	})
}