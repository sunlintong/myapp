package controllers

import (
	"myapp/modles/local"
	"myapp/modles/db"
	"encoding/json"
	"time"
	"fmt"
	"io/ioutil"
)

type ContainerRunningController struct {
	BaseController
}

type ContainerLogRequest struct {
	Container_ID string `json:"container_id`
}

func (crc *ContainerRunningController) GetRunningContainers() {
	containers, err := local.GetRunningContainers()
	l := new(db.Log)
	l.Name = "admin"
	l.Time = time.Now().Unix()
	if err != nil {
		l.Log = fmt.Sprint("get all running containers failed, %v", err)
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	} else {
		l.Log = "get all running containers succeed"
		db.InsertLog(l)
	}

	var ret [][7]string
	for _, container := range containers {
		var data [7]string
		data[0] = container.ID
		data[1] = container.Image
		data[2] = container.Command
		data[3] = crc.GetTimeString(container.Created)
		data[4] = container.Status
		data[5] = local.GetContainerPort(container)
		for _, str := range container.Names {
			data[6] += str
		}

		ret = append(ret, data)
	}
	crc.Data["json"] = ret
	crc.ServeJSON()
	crc.Finish()
}

func (crc *ContainerRunningController) GetContainerLog() {
	var req ContainerLogRequest
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(crc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.BadRequest(l)
		return
	}
	fmt.Println("req----", req)
	out, err := local.GetContainerLog(req.Container_ID)
	defer out.Close()
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	}
	msg, err := ioutil.ReadAll(out)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	}
	fmt.Println("msg----", msg)
	l.Log = "get container log success"
	db.InsertLog(l)
	crc.Success(msg)
}