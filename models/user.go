package models

import "time"

type User struct {
	Id        int64
	Name      string
	CreatedAt time.Time `xorm:"created"`
}
