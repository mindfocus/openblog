package routers

import (
	"github.com/astaxie/beego"
	"github.com/mindfocus/openblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
