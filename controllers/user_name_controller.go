package controllers

import (

)

type UserNameController struct {
	BaseController
}

func (unc *UserNameController) GetName() {
	var name string
	if unc.User.IsAdmin {
		name = "管理员 " + unc.User.Name
	} else {
		name = "普通用户 " + unc.User.Name
	}
	unc.Success(name)
}