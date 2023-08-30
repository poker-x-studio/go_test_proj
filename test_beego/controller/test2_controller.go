package controller

import "github.com/beego/beego"

type Test2Controller struct {
	beego.Controller
}

func (c *Test2Controller) Get() {
	c.Ctx.WriteString("Test2Controller.Get")
}

func (c *Test2Controller) Post() {
	c.Ctx.WriteString("Test2Controller.Post")
}
