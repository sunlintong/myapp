package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 登录注册 路由
	beego.Router("/register", &controllers.RegisterController{})

	// admin页面 路由
	beego.Router("/admin-index", &controllers.AdminIndexController{})
	beego.Router("/admin-image", &controllers.AdminImageController{})
	beego.Router("/admin-container", &controllers.AdminContainerController{})
	beego.Router("/admin-log", &controllers.AdminLogController{})
	

	// api 路由
	beego.Router("/api/user", &controllers.UserDataController{}, "get:GetUserData")
	beego.Router("/api/log", &controllers.UserLogController{}, "get:GetUserLog")

	// container 路由
	beego.Router("/api/container", &controllers.ContainerController{}, "get:GetAllContainers")
	beego.Router("/api/container/operate", &controllers.ContainerController{}, "post:OperationContainer")

	// image 路由
	beego.Router("/api/image", &controllers.ImageController{}, "get:GetAllImages")

	
}
