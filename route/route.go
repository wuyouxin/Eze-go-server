package route

import (
	"Eze-go-server/controller"
	"Eze-go-server/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

func Hub() *iris.Application {

	app := iris.Default()

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}

	app.UseGlobal(func(ctx context.Context) {
		ctx.Next()
	})

	app.DoneGlobal(func(ctx context.Context) {
	})

	mvc.Configure(app.Party("/api", crs).AllowMethods(iris.MethodOptions), func(mvc *mvc.Application) {

		mvc.Party("/user", crs).Register(service.NewUserService()).Handle(new(controller.UserController))
		mvc.Party("/menu", crs).Register(service.NewMenuService()).Handle(new(controller.MenuController))

	})

	return app

}
