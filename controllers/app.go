package controllers

import (
	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

func (c *AppController) Prepare() {
	if uid, ok := c.GetSession("uid").(int); !ok || uid <= 0 {
		c.Abort("403")
	}
}
