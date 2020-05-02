package router

import (
	"github.com/astaxie/beego"
	"github.com/personal-website-go/controller"
)

func init() {
	beego.Router("/v1/content", &controller.ContentController{}, "post:Create")
	beego.Router("/v1/content", &controller.ContentController{}, "put:Update")
	beego.Router("/v1/content/:id", &controller.ContentController{}, "get:FindById")
}

// AddNamespace for swagger
func AddNamespace() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/content",
				beego.NSInclude(
					&controller.ContentController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
