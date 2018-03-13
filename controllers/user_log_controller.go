package controllers

import (
	"github.com/golang/glog"
	"myapp/modles/db"
	"time"
)

type UserLogController struct {
	BaseController
}

func (ulc *UserLogController) GetUserLog() {
	// 生成log
	l := new(db.Log)
	l.Time = time.Now().Unix() 
	l.Name = "admin"
	l.Log = "see user's log"
	err := db.InsertLog(l)
	if err != nil {
		glog.V(3).Infoln(err)
	}

	logs, err := db.GetAllLogs()
	if err != nil {
		glog.V(3).Infoln(err)
	}
	
	var ret [][3] string
	for _, log := range logs {
		var data [3]string
		data[0] = ulc.GetTimeString(log.Time)
		data[1] = log.Name
		data[2] = log.Log
		ret = append(ret, data)
	}
	ulc.Data["json"] = ret
	glog.V(1).Infoln(ulc.Data)
	ulc.ServeJSON()
	ulc.Finish()
}