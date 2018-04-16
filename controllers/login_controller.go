package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"time"
	"myapp/types"
	"github.com/astaxie/beego/orm"
	"log"
)

type LoginController struct {
	BaseController
}

type LoginRequest struct {
	User_Name     string `json:"user_name"`
	User_Password string `json:"user_password"`
}

// 登录页面
func (lc *LoginController) Get() {
	lc.TplName = "login.html"
}

// 处理登录操作
func (lc *LoginController) Post() {
	var req LoginRequest
	l := new(db.Log)
	l.Time = time.Now().Unix()

	json.Unmarshal(lc.Ctx.Input.RequestBody, &req)
	// 这个err必须处理，因为有可能数据库没有该用户
	dbusers, err := db.GetUserByName(req.User_Name)
	log.Println("获取用户：%+v, err : v%", dbusers, err)
	// 出现这种错误是因为该用户还未注册
	if err == orm.ErrNoRows {
		l.Name = "unknown"
		l.Log = "user not found, please register first"
		err := db.InsertLog(l)
		lc.CheckErr(err)
		lc.BadRequest(l)
		lc.Redirect("/register", 302)
		return
	}
	dbuser := dbusers[0]
	if lc.Encode(req.User_Password) != dbuser.Password {
		l.Name = req.User_Name
		l.Log = "input error"
		err := db.InsertLog(l)
		lc.CheckErr(err)
		lc.BadRequest(l)
		return
	}

	// 每次登录成功都要重新设置session,避免不同用户互相影响
	user := types.User{
		Name: req.User_Name,
		IsAdmin: dbuser.UserType == "admin",
	}
	lc.DelSession("user")
	lc.SetSession("user", user)
	l.Name = user.Name
	l.Log = "login succeed"
	db.InsertLog(l)
	lc.Success(l)
}
