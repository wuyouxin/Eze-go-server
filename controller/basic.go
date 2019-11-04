package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type BasicController struct {
	Ctx     iris.Context
}

// 激活前
func (c *BasicController) BeforeActivation(mvc mvc.BeforeActivation) {
	fmt.Println("========== >>> BasicController BeforeActivation 激活前")
	mvc.Handle("GET", "/custom/custom2", "Custom")
}

// 激活后
func (c *BasicController) AfterActivation(mvc mvc.AfterActivation) {
	fmt.Println("========== >>> BasicController AfterActivation 激活后")
	mvc.Handle("GET", "/custom/cu", "Wuyou")
}

func (c *BasicController) Custom(ctx iris.Context) {
	fmt.Println("========== >>> BasicController Custom")
	_, _ = ctx.HTML("<h1>BasicController Custom</h1>")
	ctx.Next()
}

func (c *BasicController) Wuyou(ctx iris.Context) {
	fmt.Println("========== >>> BasicController Wuyou")
	_, _ = ctx.HTML("<h1>BasicController Wuyou</h1>")
	ctx.Next()
}

func (c *BasicController) Get(ctx iris.Context) {
	fmt.Println("========== >>> BasicController Get")
	_, _ = ctx.HTML("<h1>BasicController Get</h1>")
	ctx.Next()
}
