package controllers

import (
	//	"github.com/astaxie/beego"
	"myapp/modles/db"
	"time"
	"encoding/json"
)

type RegisterController struct {
	BaseController
}

type RegisterRequest struct {
	User_Name     string `json:"user_name"`
	User_Password string `json:"user_password"`
	IsAdmin bool `json:"isadmin"`
}

func (rc *RegisterController) Get() {
	rc.TplName = "register.html"
}

func (rc *RegisterController) Post() {
	var req RegisterRequest
	l := new(db.Log)
	l.Time = time.Now().Unix()
	json.Unmarshal(rc.Ctx.Input.RequestBody, &req)
	if req.IsAdmin && db.HasAdmin() {
		l.Name = "unknown"
		l.Log = "can't register as admin, system has existed an admin"
		err := db.InsertLog(l)
		rc.CheckErr(err)
		rc.BadRequest(l)
		return
	}
	dbuser, err := db.GetUserByName(req.User_Name)
	rc.CheckErr(err)
	// 找到该用户，说明用户名重复了
	if dbuser != nil {
		l.Name = "unknown"
		l.Log = "account name repeat, please register another name"
		err := db.InsertLog(l)
		rc.CheckErr(err)
		rc.BadRequest(l)
		return
	}
	user := &db.User{
		Name: req.User_Name,
		Password: rc.Encode(req.User_Password),
		RegisterTime: time.Now().Unix(),
	}
	if req.IsAdmin {
		user.UserType = "admin"
	} else {
		user.UserType = "general"
	}
	err = db.CreateUser(user)
	rc.CheckErr(err)
	l.Name = user.Name
	l.Log = "register succeed"
	db.InsertLog(l)
	rc.Redirect("/login", 302)
}
