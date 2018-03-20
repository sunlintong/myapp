package controllers

import (
	"encoding/json"
	"fmt"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
)

// 时间类型
const (
	stopEvent   = "stop"   // 停止容器
	startEvent  = "start"  // 启动容器
	removeEvent = "remove" // 删除容器
	runEvent    = "run"    // 运行一个新的容器
)

type ContainerController struct {
	BaseController
}

type ContainerRequest struct {
	Container_ID string `json:"container_id"`
	Event_Type   string `json:"event_type"`
}

func (cc *ContainerController) GetAllContainers() {
	containers, err := local.GetAllContainers()
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()
	if err != nil {
		l.Log = fmt.Sprint("get all containers failed, %v", err)
	} else {
		l.Log = "get all containers succeed"
	}
	_ = db.InsertLog(l)

	var ret [][7]string
	for _, container := range containers {
		var data [7]string
		data[0] = string([]byte(container.ID)[:20])
		data[1] = container.Image
		data[2] = container.Command
		data[3] = cc.GetTimeString(container.Created)
		data[4] = container.Status
		data[5] = local.GetContainerPort(container)
		for _, str := range container.Names {
			data[6] += str
		}

		ret = append(ret, data)
	}

	cc.Data["json"] = ret
	cc.ServeJSON()
	cc.Finish()
}

// 处理容器操作的所有post请求
func (cc *ContainerController) OperationContainer() {
	var req ContainerRequest
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(cc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		cc.BadRequest(l)
		return
	}

	// 事件的逻辑处理部分
	switch req.Event_Type {
	case stopEvent:
		err = local.StopContainer(req.Container_ID)
	case startEvent:
	case removeEvent:
	case runEvent:

	default:
		err = fmt.Errorf("unknown event %v", req.Event_Type)
	}

	if err == nil {
		l.Log = fmt.Sprintf("s% container event handle success", req.Event_Type)
	}else {
		l.Log = err.Error()
	}
	db.InsertLog(l)
	cc.Success(l)
	return

}
