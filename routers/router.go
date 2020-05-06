package routers

import (
	"pirain_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.UserController{},"*:Register")
	beego.Router("/", &controllers.UserController{},"*:Get")
	beego.Router("/charge", &controllers.UserController{},"*:UserCharge")
	beego.Router("/tran", &controllers.UserController{},"*:Trans")
}
