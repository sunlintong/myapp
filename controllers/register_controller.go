package controllers

import (
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

	//输入不完整
	if name == "" || password == "" || userType == ""{		
		rc.TplName = "inputInvalid.tpl"
		glog.Infof("a user register wrong,his info :%v",user)
	//用户名重复
	}else if u, err := db.GetUserByName(name); u != nil && err == nil{
		glog.Errorln("user input name repeat")
		rc.TplName = "nameRepeat.tpl"
	//注册信息合法，往数据库添加
	}else if err := db.CreateUser(user); err != nil {
		glog.Errorln(err)
	}else{
		glog.Infof("a user register success,his info :%v",user)
		rc.TplName = "signupSuccess.tpl"
	}	
}

