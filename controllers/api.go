package controllers

type ApiController struct {
	AppController
}

// @Title api
// @Description 注册
// @Success 200 {string} "ok"
// @Failure 400 register info not enough
// @Failure 500 register error
// @router /list [post]
func (c *ApiController) List() {
	c.ServeJSON()
}
