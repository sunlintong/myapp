package controllers

import (
)

// AdminController ...
type AdminLogController struct {
	BaseController
}

// Get ...
func (ac *AdminLogController) Get() {
	ac.TplName = "admin-log.html"
}

