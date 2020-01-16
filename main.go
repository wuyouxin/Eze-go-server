package main

import (
	"Eze-go-server/route"
	"github.com/kataras/iris/v12"
)

func main() {
	app := route.Hub()
	app.Logger().SetLevel("debug")
	_ = app.Run(iris.Addr(":9000"))
}
