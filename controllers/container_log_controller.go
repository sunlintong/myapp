package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"log"

	"github.com/docker/docker/api/types"
)

type ContainerRunningController struct {
	BaseController
}

type ContainerLogRequest struct {
	Container_ID string `json:"container_id`
	ShowStdout   bool   `json:"showstdout"`
	ShowStderr   bool   `json:"showstderr"`
	Timestamps   bool   `json:"timestamps"`
	Details      bool   `json:"details"`
	Tail         string `json:"tail"`
}

// get
func (crc *ContainerRunningController) GetRunningContainers() {
	ids, err := db.GetContainerIdsByUser(crc.User)
	crc.CheckErr(err)
	containers, err := local.GetRunningContainersGrepIds(ids)
	l := new(db.Log)
	l.Name = crc.User.Name
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

// post
func (crc *ContainerRunningController) GetContainerLog() {
	var req ContainerLogRequest
	l := new(db.Log)
	l.Name = crc.User.Name
	l.Time = time.Now().Unix()

	err := json.Unmarshal(crc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	}
	log.Println("req----", req)
	options := types.ContainerLogsOptions{}
	options.ShowStdout = req.ShowStdout
	options.ShowStderr = req.ShowStderr
	options.Timestamps = req.Timestamps
	options.Details = req.Details
	options.Tail = req.Tail
	out, err := local.GetContainerLog(req.Container_ID, options)
	defer out.Close()
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	}
	msg, err := ioutil.ReadAll(out)
	str := string(msg)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		crc.ServiceError(l)
		return
	}
	log.Println("msg----", str)
	l.Log = "get container log success"
	db.InsertLog(l)
	crc.Success(str)
}
