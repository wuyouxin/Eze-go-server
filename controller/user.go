package controller

import (
	"Eze-go-server/dao"
	"github.com/kataras/iris"
)

type UserController interface {
	WuUser()
}

func WuUser(ctx iris.Context) {
	wu := dao.GetAll()
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 wu",
		"data": wu,
	})
}

func YouUser(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 you",
	})
}
