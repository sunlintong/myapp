package controllers

// AdminController ...
type AdminContainerLogController struct {
	BaseController
}

// Get ...
func (ac *AdminContainerLogController) Get() {
	ac.TplName = "admin-container-log.html"
}

