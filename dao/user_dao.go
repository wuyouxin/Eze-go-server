package dao

import (
	"Eze-go-server/datasource"
	"Eze-go-server/models"
)

func GetAll() []models.User {
	engine := datasource.DB
	datalist := make([]models.User, 0)
	_ = engine.Desc("id").Find(&datalist)
	return datalist
}

func Insert() {
	engine := datasource.DB
	user := &models.User{Password: "dsddd"}
	_, _ = engine.Insert(user)
}
