package controllers

import (
	"fmt"
//	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"myapp/modles/db"
	"time"
)

type RegisterController struct {
	BaseController
}

func (rc *RegisterController) Get() {
	glog.Infoln("a user went into the login and register web UI")
	rc.TplName = "register.html"
}

func (rc *RegisterController) Post() {
	inputs := rc.Input()	
	name := inputs.Get("name")
	password := rc.Encode(inputs.Get("password"))
	userType := inputs.Get("userType")

	user := new(db.User)
	user.Name = name
	user.Password = password
	user.UserType = userType
	user.RegisterTime = time.Now().Unix()

	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	//输入不完整
	if user.Ileagel() {		
		rc.TplName = "inputInvalid.tpl"
		l.Log = fmt.Sprintf("a user register wrong,his info :%v", user)
		db.InsertLog(l)
		glog.Infof("a user register wrong,his info :%v", user)
	//用户名重复
	}else if u, err := db.GetUserByName(name); u != nil && err == nil{
		l.Log = fmt.Sprintf("a user register name repeat,his info :%v", user)
		db.InsertLog(l)
		rc.TplName = "nameRepeat.tpl"
	//注册信息合法，往数据库添加
	}else if err := db.CreateUser(user); err != nil {
		glog.Errorln(err)
	}else{
		l.Log = fmt.Sprintf("a user register success :%v", user)
		db.InsertLog(l)
		glog.Infof("a user register success,his info :%v", user)
		rc.Redirect("/main", 302)
	}
}

