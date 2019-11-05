package service

import (
	"Eze-go-server/dao"
	"Eze-go-server/datasource"
	"Eze-go-server/models"
)

type UserService interface {
	GetAll() []models.User
}

type userService struct {
	userDao *dao.UserDao
}

func NewUserService() *userService {
	return &userService{userDao: dao.NewUserDao(datasource.MysqlEngine())}
}

func (u *userService) GetAll() []models.User {
	return u.userDao.GetAll()
}
