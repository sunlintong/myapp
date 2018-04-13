package controllers

import (
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"encoding/json"

)

type ContainerCreateController struct {
	BaseController
}

type ContainerCreateRequest struct {
	Image_Name string `json:"image_name"`
	Container_Cmd string `json:"container_cmd"`
}

func (ccc *ContainerCreateController) CreateContainer() {
	var req ContainerCreateRequest
	l := new(db.Log)
	l.Name = "admin"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(ccc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.BadRequest(l)
		return
	}

	body, err := local.CreateContainer(req.Image_Name, []string{req.Container_Cmd})
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ccc.BadRequest(l.Log)
		return
	}
	l.Log = "create container succeed"
	db.InsertLog(l)
	ccc.Success(body)
}