package controllers

import (
	"myapp/modles/local"
	"myapp/modles/db"
	"time"
)

type ImagePullController struct {
	BaseController
}

func (ipc *ImagePullController) PullImage() {
	inputs := ipc.Input()
	imageName := inputs.Get("imageName")
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	if imageName == "" {
		l.Log = "you didn't input image name"
		db.InsertLog(l)
		ipc.BadRequest(l)
		return
	}

	out, err := local.PullImage(imageName)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.BadRequest(l)
		return
	}
	var p []byte
	_, err = out.Read(p)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.BadRequest(l)
		return
	}
	out.Close()
	l.Log = string(p)
	db.InsertLog(l)
	ipc.Success(l)	
}
