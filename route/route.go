package route

import (
	"Eze-go-server/controller/admin"
	"Eze-go-server/controller/party"
	"Eze-go-server/controller/user"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Hub() *iris.Application {

	app := iris.Default()

	// 函数式
	app.PartyFunc("/party", func(ctx iris.Party) {
		ctx.Get("/wu", party.Wu)
		ctx.Get("/you", party.You)
	})

	// v1版本
	v1 := app.Party("/v1")
	v1.Use(func(ctx context.Context) {
		ctx.Next()
	})
	{
		v1.Get("/wu", admin.Wu)
		v1.Get("/you", admin.You)
	}

	// v2版本
	v2 := app.Party("/v2")
	v2.Use(func(ctx context.Context) {
		ctx.Next()
	})
	{
		v2.Get("/wu", user.Wu)
		v2.Get("/you", user.You)
	}

	return app

}
