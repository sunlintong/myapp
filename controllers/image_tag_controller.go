package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"fmt"
)

type ImageTagController struct {
	BaseController
}

type ImageTagRequest struct {
	Source_Name string `json:"source_name"`
	Target_Name string `json:"target_name"`
	Event_Type   string `json:"event_type"`
}

const (
	DeleteTag = "deteletag"
	AddTag    = "addtag"
)

func (itc *ImageTagController) OperateTag() {
	var req ImageTagRequest
	l := new(db.Log)
	l.Name = "admin"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(itc.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		itc.BadRequest(l)
		return
	}

	switch req.Event_Type {
	case DeleteTag:
		// 因docker api不全未能实现删除tag
	case AddTag:
		if req.Source_Name == "未命名" {
			l.Log = "can't add tag to a no name image"
			db.InsertLog(l)
			itc.BadRequest("不能给未命名的镜像打标签")
			return
		}
		err = local.AddTag(req.Source_Name, req.Target_Name)
	default:
		err = fmt.Errorf("unknown event %v", req.Event_Type)
	}

	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		itc.ServiceError(l)
		return
	}else {
		l.Log = "add tag succeed"
		db.InsertLog(l)
		itc.Success(l)
	}
}
