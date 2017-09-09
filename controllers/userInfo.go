package controllers

import (
	"os"

	"strconv"

	"github.com/TalentFeng/GeekChat-server/models"
	"github.com/TalentFeng/GeekChat-server/tools"
	"github.com/astaxie/beego"
)

type UserInfoController struct {
	AppController
}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} models.UserInfo
// @Failure 400 "失败"
// @Failure 500 服务器错误
func (c *UserInfoController) Get() {
	uid := c.GetSession("uid").(int)
	var user = models.UserInfo{Uid: uid}
	db := models.GetDb()
	defer db.Close()
	db.Find(user, user)
	c.Data["json"] = user
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
	mime := h.Header.Get("Content-Type")
	if mime[0:5] != "image" {
		beego.Critical(h.Header.Get("Content-Type"))
		c.Abort("400")
	}

	var (
		db  = models.GetDb()
		uid = c.GetSession("uid").(int)
		dir = "static/upload/" + strconv.Itoa(uid)
	)
	defer db.Close()
	os.MkdirAll(dir, 0777)
	w, _ := os.OpenFile(dir+"/"+tools.Md5sum(h.Filename)+"."+mime[6:], os.O_CREATE|os.O_RDWR, 0777)
	err = tools.ImageScale(f, w, 1280, 720)
	if err != nil {
		beego.Critical(err)
		c.Abort("500")
	}
	err = db.Save(&models.UserInfo{Uid: uid, Avatar: tools.Md5sum(h.Filename) + "." + mime[6:]}).Error
	if err != nil {
		beego.Critical(err)
		c.Abort("500")
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}
