package controllers

// AdminController ...
type AdminLogController struct {
	BaseController
}

// Get ...
func (ac *AdminLogController) Get() {
	ac.TplName = "log.html"
}

