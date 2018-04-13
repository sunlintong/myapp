package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"fmt"
)

type ContainerCommitController struct {
	BaseController
}

type ContainerCommitRequest struct {
	Container_ID string `json:"container_id"`
	Image_Name string `json:"image_name"`
}

func (ccc *ContainerCommitController) Commit() {
	var req ContainerCommitRequest
	l := new(db.Log)
	l.Name = "admin"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(ccc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.ServiceError(l)
		return
	}
	fmt.Println("req----", req)

	image_id, err := local.CommitContainer(req.Container_ID, req.Image_Name)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.ServiceError(l)
		return
	}
	fmt.Println("NEW IMAGE :", image_id)
	l.Log = fmt.Sprintf("commit container succeed, new image id : %s", image_id)
	db.InsertLog(l)
	ccc.Success(image_id)
}