package controllers

import (
	"fmt"
	"myapp/modles/db"
	"github.com/golang/glog"
	"time"
	"log"
)

type UserDataController struct {
	BaseController
}

// 获取用户信息
func (udc *UserDataController) GetUserData() {
	users, err := db.GetAllUsers()
	if err != nil {
		glog.V(3).Infoln(err)
	}

	// 生成log
	log := new(db.Log)
	log.Time = time.Now().Unix() 
	log.Name = "admin"
	log.Log = "see user's infomation"
	err = db.InsertLog(log)
	if err != nil {
		glog.V(3).Infoln(err)
	}

	var ret [][3]string

	for _, user := range users {
		var data [3]string
		data[0] = user.Name
		data[1] = user.UserType
		data[2] = udc.GetTimeString(user.RegisterTime)
		ret = append(ret, data)
		fmt.Println(user.ID, data)
	}

	udc.Data["json"] = ret
	glog.V(1).Infoln(udc.Data)
	udc.ServeJSON()
	udc.Finish()
}

func (udc *UserDataController) Logout() {
	l := new(db.Log)
	l.Time = time.Now().Unix()
	l.Name = udc.User.Name
	udc.DelSession("user")
	l.Log = fmt.Sprintf("user %s logout", udc.User.Name)
	log.Printf(l.Log)
	err := db.InsertLog(l)
	if err != nil {
		udc.ServiceError(l)
	}
	udc.Success("l")
}
