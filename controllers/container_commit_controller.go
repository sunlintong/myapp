package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
)

type ContainerCommitController struct {
	BaseController
}

type ContainerCommitRequest struct {
	Container_ID string `json:"container_id"`
	Image_Name   string `json:"image_name"`
}

func (ccc *ContainerCommitController) Commit() {
	var req ContainerCommitRequest
	l := new(db.Log)
	l.Name = ccc.User.Name
	l.Time = time.Now().Unix()

	err := json.Unmarshal(ccc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.ServiceError(l)
		return
	}
	log.Println("req----", req)

	image_id, err := local.CommitContainer(req.Container_ID, req.Image_Name)
	dbimage := &db.Image{
		Image_ID: image_id,
		Group:    ccc.User.Name,
	}
	db.InsertImage(dbimage)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.ServiceError(l)
		return
	}
	log.Println("NEW IMAGE :", image_id)
	l.Log = fmt.Sprintf("commit container succeed, new image id : %s", image_id)
	db.InsertLog(l)
	ccc.Success(image_id)
}
