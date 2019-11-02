package dao

import (
	"Eze-go-server/datasource"
	"Eze-go-server/models"
)

func GetAll() []models.User {
	engine := datasource.DB
	datalist := make([]models.User, 0)
	err := engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
