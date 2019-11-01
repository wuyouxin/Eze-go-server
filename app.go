package main

import (
	_ "Eze-go-server/engine"
	"Eze-go-server/route"
	"github.com/kataras/iris"
)

func main() {
	_ = route.Hub().Run(iris.Addr(":9000"))
}
