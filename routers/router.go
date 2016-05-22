package routers

import (
	"github.com/astaxie/beego"
	"pro_monitor/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.AppLogController{})
}
