package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 登录注册 路由
	beego.Router("/register", &controllers.RegisterController{})

	// 页面 路由
	beego.Router("/admin-index", &controllers.AdminIndexController{})
	beego.Router("/admin-image", &controllers.AdminImageController{})
	beego.Router("/admin-container", &controllers.AdminContainerController{})
	beego.Router("/admin-container-log", &controllers.AdminContainerLogController{})
	beego.Router("/admin-log", &controllers.AdminLogController{})

	

	// api 路由
	beego.Router("/api/user", &controllers.UserDataController{}, "get:GetUserData")

	beego.Router("/api/log", &controllers.UserLogController{}, "post:GetUserLog")

	beego.Router("/api/container", &controllers.ContainerController{}, "get:GetAllContainers")
	beego.Router("/api/container/operate", &controllers.ContainerController{}, "post:OperateContainer")
	beego.Router("/api/container/running", &controllers.ContainerRunningController{}, "get:GetRunningContainers")
	beego.Router("/api/container/log", &controllers.ContainerRunningController{}, "post:GetContainerLog")
	beego.Router("/api/container/commit", &controllers.ContainerCommitController{}, "post:Commit")
	beego.Router("/api/container/create", &controllers.ContainerCreateController{}, "post:CreateContainer")

	beego.Router("/api/image", &controllers.ImageController{}, "get:GetAllImages")
	beego.Router("/api/image/operate", &controllers.ImageController{}, "post:OperateImage")
	beego.Router("/api/image/pull", &controllers.ImagePullController{}, "post:PullImage")
	beego.Router("/api/image/tag", &controllers.ImageTagController{}, "post:OperateTag")
	
}
