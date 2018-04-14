package controllers

import (
	"myapp/modles/db"
	"time"
	"fmt"
)

type UserLogController struct {
	BaseController
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
	var ret [][3] string
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
	name := ulc.GetString("user_name", "")
	fmt.Println("nnnnn", name)
	ulc.SetSession("user_name", name)
}