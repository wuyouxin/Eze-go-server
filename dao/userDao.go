package dao

import (
	"Eze-go-server/models"
	"github.com/go-xorm/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine: engine}
}


func (u *UserDao) GetAll() []models.User {
	datalist := make([]models.User, 0)
	_ = u.engine.Desc("id").Find(&datalist)
	return datalist
}
