package controller

import (
	"Eze-go-server/dao"
	"github.com/kataras/iris"
)

func GetUserAll(ctx iris.Context) {
	data := dao.GetAll()
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 wu",
		"data": data,
	})
}

func YouUser(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"msg":  "v2 you",
	})
}
