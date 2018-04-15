package controllers

import (
	"fmt"
	"myapp/modles/db"
	"time"
	"log"
)

type UserDataController struct {
	BaseController
}

// 获取用户信息
func (udc *UserDataController) GetUserData() {
	var users []*db.User
	var err error
	l := new(db.Log)
	l.Time = time.Now().Unix() 
	l.Name = udc.User.Name
	// 非admin只能看到自己
	if !udc.User.IsAdmin {
		users, err = db.GetUserByName(udc.User.Name)
	}else {
		users, err = db.GetAllUsers()
	}
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		udc.BadRequest(l)
		return
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

	udc.Success(ret)
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
