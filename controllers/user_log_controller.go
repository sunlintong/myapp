package controllers

import (
	"myapp/modles/db"
	"time"
)

type UserLogController struct {
	BaseController
}

type UserLogRequest struct{
	User_Name string `json:"user_name"`
}

// post 根据用户名获取log
// 若用户名为kong，则获取所有log
func (ulc *UserLogController) GetUserLog() {
	l := new(db.Log)
	l.Time = time.Now().Unix()
	l.Name = "admin"
	// 若未传name，则说明是admin在看所有用户日志
	name := ulc.GetString("user_name", "")
	var err error
	var logs []*db.Log
	if name == "" {
		logs, err = db.GetAllLogs()
	} else {
		logs, err = db.GetLogsByUser(name)
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