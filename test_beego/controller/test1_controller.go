package controller

import "github.com/beego/beego"

type Test1Controller struct {
	beego.Controller
}

func (c *Test1Controller) Get() {
	c.Ctx.WriteString("Test1Controller.Get")
}

func (c *Test1Controller) Post() {
	c.Ctx.WriteString("Test1Controller.Post")
}
