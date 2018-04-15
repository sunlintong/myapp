package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"time"
	"log"
)

type UserLogController struct {
	BaseController
}

type UserLogRequest struct {
	User_Name string `json:"user_name"`
}

// get
func (ulc *UserLogController) GetUserLog() {
	l := new(db.Log)
	l.Time = time.Now().Unix()
	l.Name = "admin"

	var err error
	var logs []*db.Log
	// 获取请求用户名，若没有，则设为""
	name := ulc.GetSession("user_name")
	if name == nil {
		name = ""
	}
	ulc.DelSession("user_name")

	// 先判断是否允许当前用户权限
	// 若当前用户不是admin确请求查看其它用户日志，不允许
	if !ulc.User.IsAdmin && name != ulc.User.Name {
		log.Printf("current_user:%+v, request user name:%s", ulc.User, name)
		ulc.BadRequest("cant request others log")
		return
	}

	if name == "" {
		logs, err = db.GetAllLogs()
	} else {
		logs, err = db.GetLogsByUser(name.(string))
	}
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ulc.ServiceError(l.Log)
		return
	}
	var ret [][3]string
	for _, log := range logs {
		var data [3]string
		data[0] = "[" + ulc.GetTimeString(log.Time) + "]"
		data[1] = "[" + log.Name + "]"
		data[2] = "---" + log.Log
		ret = append(ret, data)
	}
	l.Log = "get log succeed"
	db.InsertLog(l)
	ulc.Success(ret)
}

// post 设置用户名
func (ulc *UserLogController) SetUserSession() {
	// 若未传name，则说明是admin在看所有用户日志
	var req UserLogRequest
	json.Unmarshal(ulc.Ctx.Input.RequestBody, &req)
	ulc.SetSession("user_name", req.User_Name)
	log.Println(ulc.CruSession)
	ulc.Success("set session success")
}
