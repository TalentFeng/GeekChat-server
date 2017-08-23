package controllers

import (
	"os"

	"strconv"

	"github.com/TalentFeng/GeekChat-server/models"
	"github.com/astaxie/beego"
)

type UserInfoController struct {
	AppController
}

func (c *UserInfoController) Get() {
	c.ServeJSON()
}

// @Title 设置头像
// @Description 设置头像
// @Param img query file true 头像
// @Success 200 {string} "ok"
// @Failure 400 "图片非法"
// @Failure 500 服务器错误
// @router /setAvatar [post]
func (c *UserInfoController) SetAvatar() {
	f, h, err := c.GetFile("img")
	if err != nil {
		beego.Critical(err)
		c.Abort("400")
	}
	defer f.Close()
	var (
		db      = models.GetDb()
		user, _ = c.Ctx.Input.Session("user").(models.User)
		dir     = "static/upload/" + strconv.Itoa(c.GetSession("uid").(int))
	)
	defer db.Close()
	os.MkdirAll(dir, 777)
	beego.Info(dir)
	err = c.SaveToFile("img", dir+"/"+h.Filename)
	if err != nil {
		beego.Critical(err)
		c.Abort("500")
	}
	err = db.Save(&models.UserInfo{Uid: user.Id, Avatar: h.Filename}).Error
	if err != nil {
		beego.Critical(err)
		c.Abort("500")
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}
