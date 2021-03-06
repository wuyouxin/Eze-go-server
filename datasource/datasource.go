package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"xorm.io/core"
)

func MysqlEngine() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/eze?charset=utf8")

	if err != nil {
		golog.Fatalf("初始化数据库连接失败!! %s", err)
	}

	engine.ShowSQL(true)

	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "sys_"))

	return engine
}
