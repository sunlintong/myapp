package controllers

import (
)

// AdminController ...
type AdminContainerController struct {
	BaseController
}

// Get ...
func (ac *AdminContainerController) Get() {
	ac.TplName = "container.html"
}

