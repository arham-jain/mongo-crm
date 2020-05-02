package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"] = append(beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"] = append(beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"] = append(beego.GlobalControllerRouter["github.com/personal-website-go/controller:ContentController"],
        beego.ControllerComments{
            Method: "FindById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
