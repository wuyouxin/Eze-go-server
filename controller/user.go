package controller

import (
	"Eze-go-server/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (u *UserController) Get() mvc.Result {
	data := u.Service.GetAll()[1]
	return mvc.Response{
		Object: map[string]interface{}{
			"status": iris.StatusOK,
			"data":   data,
		},
	}
}
