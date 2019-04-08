package routers

import (
	"gocms/controllers"
	"github.com/astaxie/beego"
	"fmt"
	"gocms/service"
	"github.com/astaxie/beego/context"
)

var navList = func(context *context.Context) {
	fmt.Println("拦截")
	context.Input.SetData("Categorys",service.GetCategorys())
}

func init() {
	beego.Include(&controllers.NewsController{})
	beego.Router("/", &controllers.MainController{})

	beego.InsertFilter("/:nav/",beego.BeforeExec,navList)
	beego.InsertFilter("/:nav/:id",beego.BeforeExec,navList)
}
