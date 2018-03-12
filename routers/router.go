package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/admin-index", &controllers.AdminIndexController{})


	beego.Router("/user-data", &controllers.UserDataController{}, "get:GetUsers")
}
