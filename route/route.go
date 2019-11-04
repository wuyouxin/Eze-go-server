package route

import (
	"Eze-go-server/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
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
		v1.Get("/wu", hero.Handler(controller.GetUserAll))
		v1.Get("/you", controller.YouUser)
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

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}

	crsv := app.Party("/api/crs", crs).AllowMethods(iris.MethodOptions)
	{
		crsv.Post("/main", func(ctx context.Context) {
			ctx.Application().Logger().Info("======   >>> crsv")
		})

		crsv.Get("/home", func(ctx context.Context) {
			ctx.Application().Logger().Info("======   >>> crsv")
		})

		crsv.Delete("/home", func(ctx context.Context) {
			ctx.Application().Logger().Info("======   >>> crsv")
		})

		crsv.Put("/home", func(ctx context.Context) {
			ctx.Application().Logger().Info("======   >>> crsv")
		})

	}

	app.UseGlobal(func(ctx context.Context) {
		_, _ = ctx.HTML("<h1>Use Global</h1>")
		ctx.Next()
	})

	app.DoneGlobal(func(ctx context.Context) {
		_, _ = ctx.HTML("<h1>Done Global</h1>")
	})

	mvc.Configure(app.Party("/sub", crs).AllowMethods(iris.MethodOptions)).Handle(new(controller.BasicController))

	mvc.Configure(app.Party("/basic"), func(mvc *mvc.Application) {
		mvc.Handle(new(controller.BasicController))
	})

	return app

}
