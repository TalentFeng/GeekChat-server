package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:ApiController"] = append(beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:ApiController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/TalentFeng/GeekChat-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
