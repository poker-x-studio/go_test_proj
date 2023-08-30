package controller

import "github.com/beego/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("MainController.Get")
}

func (c *MainController) Post() {
	c.Ctx.WriteString("MainController.Post")
}
