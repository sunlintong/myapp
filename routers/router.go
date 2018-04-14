package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 路由过滤器
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uname").(string)
	// 没有session说明用户还未登录，这时如果请求其他页面，则跳转至登录页面
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "login")
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	// 登录注册 路由
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})

	// 页面 路由
	beego.Router("/admin-index", &controllers.AdminIndexController{})
	beego.Router("/admin-image", &controllers.AdminImageController{})
	beego.Router("/admin-container", &controllers.AdminContainerController{})
	beego.Router("/admin-container-log", &controllers.AdminContainerLogController{})
	beego.Router("/admin-log", &controllers.AdminLogController{})

	

	// api 路由
	beego.Router("/api/user", &controllers.UserDataController{}, "get:GetUserData")

	beego.Router("/api/log", &controllers.UserLogController{}, "get:GetUserLog")
	beego.Router("/api/log", &controllers.UserLogController{}, "post:SetUserSession")

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
