package controllers

import (
	"fmt"
	"myapp/modles/local"
	"myapp/modles/db"
	"time"
	"io/ioutil"
)

type ImagePullController struct {
	BaseController
}

func (ipc *ImagePullController) PullImage() {
	imageName := ipc.GetString("imageName")
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	if imageName == "" {
		l.Log = "you didn't input image name"
		err := db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		fmt.Println(imageName, l)
		return
	}

	out, err := local.PullImage(imageName)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		fmt.Println(imageName, l)
		return
	}

	msg, err := ioutil.ReadAll(out)
	defer out.Close()
	fmt.Println(msg)

	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		fmt.Println(imageName, l)
		return
	}
	// 只保存前250
	l.Log = string(msg[:250])
	db.InsertLog(l)
	ipc.CheckErr(err)
	fmt.Println(imageName, l)
	ipc.Success(l)	
}
