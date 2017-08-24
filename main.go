package main

import (
	"github.com/TalentFeng/GeekChat-server/models"
	_ "github.com/TalentFeng/GeekChat-server/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		models.AutoMigrate()
	}
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.ServerName = "laotanyu"
	beego.Run()
}
