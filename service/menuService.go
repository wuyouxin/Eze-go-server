package service

import (
	"Eze-go-server/dao"
	"Eze-go-server/datasource"
	"Eze-go-server/models"
)

type MenuService interface {
	GetAll() []models.Menu
}

type menuService struct {
	menuDao *dao.MenuDao
}

func NewMenuService() *menuService {
	return &menuService{menuDao: dao.NewMenuDao(datasource.MysqlEngine())}
}

func (m *menuService) GetAll() []models.Menu {
	return m.menuDao.GetAll()
}
