package controllers

import (
	"myapp/modles/local"
	"myapp/modles/db"
	"time"
	"fmt"
)

type ContainerRunningController struct {
	BaseController
}

func (crc *ContainerRunningController) GetRunningContainers() {
	containers, err := local.GetRunningContainers()
	l := new(db.Log)
	l.Name = "admin"
	l.Time = time.Now().Unix()
	if err != nil {
		l.Log = fmt.Sprint("get all containers failed, %v", err)
		db.InsertLog(l)
		crc.ServiceError(l)
	} else {
		l.Log = "get all containers succeed"
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