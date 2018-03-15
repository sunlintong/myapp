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

	var ret [][4]string
	for _, container := range containers {
		var data [4]string
		data[0] = container.ID
		data[1] = container.Names[0]
		data[2] = container.Image
		data[3] = container.Status
		ret = append(ret, data)
	}

	cc.Data["json"] = ret
	cc.ServeJSON()
	cc.Finish()
}
