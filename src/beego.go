// beego.go

package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello beego get")
}

func (this *MainController) Post() {
	this.Ctx.WriteString("hello beego post")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

