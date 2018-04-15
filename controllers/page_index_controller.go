package controllers

import (
)

// AdminController ...
type AdminIndexController struct {
	BaseController
}

// Get ...
func (ac *AdminIndexController) Get() {
	ac.TplName = "index.html"
}

