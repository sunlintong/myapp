package controllers

import (
)

// AdminController ...
type AdminImageController struct {
	BaseController
}

// Get ...
func (ac *AdminImageController) Get() {
	ac.TplName = "admin-image.html"
}

