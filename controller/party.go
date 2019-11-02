package controller

import "github.com/kataras/iris"

func WuParty(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "party wu",
	})
}

func YouParty(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "party you",
	})
}
