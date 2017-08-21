package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/TalentFeng/GeekChat-server/models"
	"github.com/astaxie/beego"
    "github.com/TalentFeng/GeekChat-server/tools"
    "github.com/astaxie/beego/validation"
)

var db *gorm.DB

func init() {
	db = models.GetDb()
}

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
func (c *UserController) Register()  {
    var user models.User
    valid := new(validation.Validation)
	user.Phone = c.Ctx.Input.Query("phone")
    user.Mail = c.Ctx.Input.Query("mail")
    user.Password = tools.Password(c.Ctx.Input.Query("password"))
    if ok, err := valid.Valid(&user); !ok {
        if err != nil {
            beego.Critical(err)
        }
        for _, err := range valid.Errors {
            beego.Info([]byte(err.Message))
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
// @Success 200 {object} models.model.User
// @Param   phone     query   int false       "手机号"
// @Param   mail      query   string false       "邮箱"
// @Param   password  query   string  false       "密码"
// @Failure 400 登陆失败
// @Failure 500 服务错误
// @router /login [post]
func (c *UserController) Login()  {
    var (
        user models.User
        valid validation.Validation
    )
    user.Password = tools.Password(c.Ctx.Input.Query("password"))
    user.Phone = c.Ctx.Input.Query("phone")
    user.Mail = c.Ctx.Input.Query("mail")
    valid.Phone(user.Phone, "phone")
    if valid.HasErrors() {
        for _, err := range valid.Errors {
            if err != nil {
                c.Abort("400")
            }
        }
    }
    if err := db.Find(&user, user).Error; err != nil {
        beego.Critical(err)
        c.Abort("400")
    }
	c.Data["json"] = user
	c.ServeJSON()
}