package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type JSONDemo struct {
	Code int
	Msg  string
}


func (c *MainController) Get() {
	mystruct := &JSONDemo{0, "hello"}
	c.Data["json"] = mystruct
	c.ServeJSON()
}
