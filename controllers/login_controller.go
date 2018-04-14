package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"time"
	"myapp/types"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

type LoginRequest struct {
	User_Name     string `json:"user_name"`
	User_Password string `json:"user_password"`
	IsAdmin       bool   `json:"isadmin"`
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
	if req.IsAdmin && db.HasAdmin() {
		l.Name = "unknown"
		l.Log = "不能注册管理员，系统已有管理员"
		err := db.InsertLog(l)
		lc.CheckErr(err)
		lc.BadRequest(l)
		return
	}
	// 这个err必须处理，因为有可能数据库没有该用户
	dbuser, err := db.GetUserByName(req.User_Name)
	// 出现这种错误是因为该用户还未注册
	if err == orm.ErrNoRows {
		l.Name = "unknown"
		l.Log = "没找到该用户，请先注册"
		err := db.InsertLog(l)
		lc.CheckErr(err)
		lc.BadRequest(l)
		return
	}
	if lc.Encode(req.User_Password) != dbuser.Password {
		l.Name = req.User_Name
		l.Log = "密码输错鸟"
		err := db.InsertLog(l)
		lc.CheckErr(err)
		lc.BadRequest(l)
		return
	}

	// 每次登录成功都要重新设置session,避免不同用户互相影响
	user := types.User{
		Name: req.User_Name,
		IsAdmin: req.IsAdmin,
	}
	lc.DelSession("user")
	lc.SetSession("user", user)
	l.Name = user.Name
	l.Log = "登录成功"
	db.InsertLog(l)
	lc.Redirect("/admin-index", 302)
}
