package engine

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	Engine, _ = xorm.NewEngine("mysql", "root:@/test?charset=utf8")
}
