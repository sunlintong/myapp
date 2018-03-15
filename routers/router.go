package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/admin-index", &controllers.AdminIndexController{})


	beego.Router("/user/data", &controllers.UserDataController{}, "get:GetUserData")
	beego.Router("/user/log", &controllers.UserLogController{}, "get:GetUserLog")

	beego.Router("/container", &controllers.ContainerController{}, "get:GetAllContainers")
}
