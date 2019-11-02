package controller

import (
	"github.com/kataras/iris"
)

func WuAdmin(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v1 wu",
	})
}

func YouAdmin(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v1 you",
	})
}
