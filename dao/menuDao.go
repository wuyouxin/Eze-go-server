package dao

import (
	"Eze-go-server/models"
	"github.com/go-xorm/xorm"
)

type MenuDao struct {
	engine *xorm.Engine
}

func NewMenuDao(engine *xorm.Engine) *MenuDao {
	return &MenuDao{engine: engine}
}

func (m *MenuDao) GetAll() []models.Menu {
	datalist := make([]models.Menu, 0)
	_ = m.engine.Desc("id").Find(&datalist)
	return datalist
}
