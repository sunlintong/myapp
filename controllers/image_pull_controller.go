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

type ImagePullReq struct {
	imageName string `json:"image_name"`
}

func (ipc *ImagePullController) PullImage() {
	imageName := ipc.GetString("image_name")
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()
	fmt.Println("dsdasfef", imageName)

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
