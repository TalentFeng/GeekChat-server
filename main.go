package main

import (
	_ "github.com/TalentFeng/GeekChat-server/routers"
	_ "github.com/TalentFeng/GeekChat-server/models"
    "github.com/astaxie/beego"
	"github.com/TalentFeng/GeekChat-server/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		models.AutoMigrate()
	}
    beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    beego.BConfig.ServerName = "laotanyuV1"
	beego.Run()
}