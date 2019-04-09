package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gocms/controllers:IndexController"] = append(beego.GlobalControllerRouter["gocms/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gocms/controllers:NewsController"] = append(beego.GlobalControllerRouter["gocms/controllers:NewsController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/:nav/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gocms/controllers:NewsController"] = append(beego.GlobalControllerRouter["gocms/controllers:NewsController"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/:nav/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
