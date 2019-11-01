package user

import "github.com/kataras/iris"

func Wu(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 wu",
	})
}

func You(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 you",
	})
}
