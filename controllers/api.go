package controllers

type ApiController struct {
    AppController
}


// @Title api
// @Description 注册
// @Success 200 {string} "ok"
// @Param   phone     query   int true       "手机号"
// @Param   mail    query   int false       "邮箱"
// @Param   password   query   string  false       "密码"
// @Failure 400 register info not enough
// @Failure 500 register error
// @router /list [post]
func (c *ApiController) List()  {
    c.ServeJSON()
}