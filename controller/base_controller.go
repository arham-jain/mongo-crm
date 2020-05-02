package controller

import "github.com/astaxie/beego"

//BaseController base class for API controllers
type BaseController struct {
	beego.Controller
}

func (i *BaseController) serve(statusCode int, model interface{}) {
	i.Data["json"] = model
	i.Ctx.Output.Status = statusCode
	i.ServeJSON()
}
