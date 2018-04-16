package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"time"
	"log"
	"fmt"
)

type UserLogController struct {
	BaseController
}

type UserLogRequest struct {
	User_Name string `json:"user_name"`
}

// get
func (ulc *UserLogController) GetUserLog() {
	var err error
	var logs []*db.Log
	l := new(db.Log)
	l.Time = time.Now().Unix()
	l.Name = "admin"
	// 获取请求用户名，若没有，则设为""
	name := ulc.GetSession("user_name")
	log.Println("user_name:", name)
	var req_name string
	n, ok := name.(string)
	if ok {
		req_name = n
	} else {
		req_name = ""
	}
	ulc.DelSession("user_name")

	// 请求为空说明是直接get日志展示界面，根据权限进行展示
	if req_name == "" {
		if ulc.User.IsAdmin {
			logs, err = db.GetAllLogs()
		}else {
			logs, err = db.GetLogsByUser(ulc.User.Name)
			log.Println("logs:", logs)
		}
	} else {
		// 用户为admin或请求当前用户的日志时才允许
		if ulc.User.IsAdmin || name == ulc.User.Name {
		logs, err = db.GetLogsByUser(req_name)
		} else {
			err = fmt.Errorf("you can't request other's log")
		}
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
