package controllers

import (
	"github.com/TalentFeng/GeekChat-server/models"
	"github.com/TalentFeng/GeekChat-server/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

// @Title 注册
// @Description 普通用户注册
// @Success 200 {string} "ok"
// @Param   phone     query   int true       "手机号"
// @Param   mail    query   string false       "邮箱"
// @Param   password   query   string  false       "密码"
// @Failure 400 register info not enough
// @Failure 500 register error
// @router /register [post]
func (c *UserController) Register() {
	var (
		user  = new(models.User)
		db    = models.GetDb()
		valid = new(validation.Validation)
	)
	defer db.Close()
	c.ParseForm(user)
	user.Password = tools.Password(user.Password)
	if ok, err := valid.Valid(user); !ok {
		if err != nil {
			beego.Critical(err)
		}
		for _, err := range valid.Errors {
			beego.Info(err.Message)
		}
		c.Abort("400")
	}
	if err := db.Save(&user).Error; err != nil {
		beego.Critical(err)
		c.Abort("500")
	}
	c.Data["json"] = "ok"
	c.ServeJSON()

}

// @Title login
// @Description 普通用户登陆
// @Success 200 {string} "ok"
// @Param   phone     query   int false       "手机号"
// @Param   mail      query   string false       "邮箱"
// @Param   password  query   string  false       "密码"
// @Failure 400 登陆失败
// @Failure 500 服务错误
// @router /login [post]
func (c *UserController) Login() {
	var (
		user  = new(models.User)
		valid validation.Validation
		db    = models.GetDb()
	)
	defer db.Close()
	c.ParseForm(user)
	user.Password = tools.Password(user.Password)
	if err := valid.Phone(user.Phone, "phone").Error; err != nil {
		beego.Info(err)
		c.Abort("400")
	}

	if err := db.Find(&user, user).Error; err != nil {
		beego.Critical(err)
		c.Abort("400")
	}
	c.Data["json"] = "ok"
	c.SetSession("uid", user.Id)
	c.ServeJSON()
}

// @Title logout
// @Description 普通用户登出
// @Success 200 {string} "ok"
// @Failure 400 登出
// @Failure 500 服务错误
// @router /logout [get]
func (c *UserController) Logout() {
	c.DelSession("uid")
	c.Data["json"] = "ok"
	c.ServeJSON()
}
