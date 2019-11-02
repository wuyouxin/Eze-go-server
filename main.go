package main

import (
	"Eze-go-server/route"
	"github.com/kataras/iris"
)

func main() {
	_ = route.Hub().Run(iris.Addr(":9000"))
}
