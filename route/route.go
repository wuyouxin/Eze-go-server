package route

import (
	"Eze-go-server/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
)

func Hub() *iris.Application {

	app := iris.Default()

	// 函数式
	app.PartyFunc("/party", func(ctx iris.Party) {
		ctx.Get("/wu", controller.WuParty)
		ctx.Get("/you", controller.YouParty)
	})

	// user版本
	v1 := app.Party("/user")
	v1.Use(func(ctx context.Context) {
		ctx.Next()
	})
	{
		v1.Get("/wu", controller.WuUser)
		v1.Get("/you", hero.Handler(controller.YouUser))
	}

	// admin版本
	v2 := app.Party("/admin")
	v2.Use(func(ctx context.Context) {
		ctx.Next()
	})
	{
		v2.Get("/wu", controller.WuAdmin)
		v2.Get("/you", controller.WuAdmin)
	}

	return app

}
