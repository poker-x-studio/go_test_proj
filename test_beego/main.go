package main

import (
	"test_beego/controller"

	"github.com/beego/beego"
)

func main() {
	init_router()
	beego.Run()
}

func init_router() {
	beego.Router("/", &controller.MainController{}, "GET:Get")

	beego.Router("/test/a/a/b", &controller.Test2Controller{}, "GET:Get")
	beego.Router("/test", &controller.Test2Controller{}, "POST:Post")

	beego.Router("/test", &controller.Test1Controller{}, "GET:Get")
	beego.Router("/test", &controller.Test1Controller{}, "POST:Post")
}
