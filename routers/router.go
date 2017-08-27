package routers

import (
	"miclowan.ru/beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
