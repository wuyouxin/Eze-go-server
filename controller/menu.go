package controller

import (
	"Eze-go-server/service"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/v12"
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
