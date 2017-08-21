// @APIVersion 1.0.0
// @Title geek chat API
// @Description geek chat web api.
// @Contact talentlqf@gmail.com
package routers

import (
	"github.com/TalentFeng/GeekChat-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/api",
			beego.NSInclude(&controllers.ApiController{}),
		),
	)
	beego.AddNamespace(ns)
}
