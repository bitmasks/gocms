package routers

import (
	"gocms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.NewsController{})
	beego.Router("/", &controllers.MainController{})
}
