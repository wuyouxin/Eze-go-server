package party

import "github.com/kataras/iris"

func Wu(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "party wu",
	})
}

func You(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "party you",
	})
}
