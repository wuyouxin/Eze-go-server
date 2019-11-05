package controller

import (
	"Eze-go-server/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MenuController struct {
	Ctx     iris.Context
	Service service.MenuService
}

func (m *MenuController) Get() mvc.Result {
	data := m.Service.GetAll()
	return mvc.Response{
		Object: map[string]interface{}{
			"status": iris.StatusOK,
			"data":   data,
		},
	}
}
