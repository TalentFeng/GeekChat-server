package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {

	c.Data["Website"] = "beego33.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["_xsrf"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "index.html"
}