package controllers

import (
	"fmt"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
)

type ContainerController struct {
	BaseController
}

func (cc *ContainerController) GetAllContainers() {
	containers, err := local.GetRunningContainers()
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()
	if err != nil {
		l.Log = fmt.Sprint("get all containers failed, %v", err)
	}else {
		l.Log = "get all containers succeed"
	}
	_ = db.InsertLog(l)
	
	var ret [][7]string
	for _, container := range containers {
		var data [7]string
		data[0] = container.ID
		data[1] = container.Image
		data[2] = container.Command
		data[3] = cc.GetTimeString(container.Created)
		data[4] = container.Status
		data[5] = fmt.Sprint(container.Ports)
		for _, str := range container.Names {
			data[6] += str
		}

		
		ret = append(ret, data)
	}

	cc.Data["json"] = ret
	cc.ServeJSON()
	cc.Finish()
}
