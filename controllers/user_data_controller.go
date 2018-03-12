package controllers

import (
	"myapp/modles/db"

	"github.com/golang/glog"
)

type UserDataController struct {
	BaseController
}

// 获取用户信息
func (udc *UserDataController) GetUsers() {
	users, err := db.GetAllUsers()
	if err != nil {
		glog.V(3).Infoln(err)
	}
	glog.V(1).Infoln(users)

	var ret [][3]string

	for _, user := range users {
		var data [3]string
		data[0] = user.Name
		data[1] = user.UserType
		data[2] = user.GetTime()
		ret = append(ret, data)
	}

	udc.Data["json"] = ret
	glog.V(1).Infoln(udc.Data)
	udc.ServeJSON()
	udc.Finish()
}
