package controllers

import (
	"log"
	"myapp/modles/local"
	"myapp/modles/db"
	"time"
	"encoding/json"
	"io/ioutil"
)

type ImagePullController struct {
	BaseController
}

type ImagePullReq struct {
	Image_Name string `json:"image_name"`
}

func (ipc *ImagePullController) PullImage() {
	var req ImagePullReq
	l := new(db.Log)
	l.Name = ipc.User.Name
	l.Time = time.Now().Unix()
	err := json.Unmarshal(ipc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		return
	}
	imageName := req.Image_Name
	if imageName == "" {
		l.Log = "you didn't input image name"
		err := db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		log.Println(imageName, l)
		return
	}

	out, err := local.PullImage(imageName)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		log.Println(imageName, l)
		return
	}

	msg, err := ioutil.ReadAll(out)
	str := string(msg)
	defer out.Close()
	log.Println(str)

	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ipc.CheckErr(err)
		ipc.BadRequest(l)
		log.Println(imageName, l)
		return
	}
	// 只保存前250
	l.Log = string(msg[:250])
	db.InsertLog(l)
	ipc.CheckErr(err)
	log.Println(imageName, l)
	ipc.Success(str)	
}
