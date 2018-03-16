package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 页面路由
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/admin-index", &controllers.AdminIndexController{})
	beego.Router("/admin-image", &controllers.AdminImageController{})
	beego.Router("/admin-container", &controllers.AdminContainerController{})


	// 用户api路由
	beego.Router("/user/data", &controllers.UserDataController{}, "get:GetUserData")
	beego.Router("/user/log", &controllers.UserLogController{}, "get:GetUserLog")

	// docker api 路由
	beego.Router("/container", &controllers.ContainerController{}, "get:GetAllContainers")
	beego.Router("/image", &controllers.ImageController{}, "get:GetAllImages")
}
