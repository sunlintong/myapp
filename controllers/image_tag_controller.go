package controllers

import (
	"encoding/json"
	"myapp/modles/db"
	"time"
)

type ImageTagController struct {
	BaseController
}

type ImageTagRequest struct {
	Image_ID   string `json:"image_id'`
	Image_Name string `json:"image_name"`
	EventType  string `json:"event_type"`
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

	switch req.EventType {
	case DeleteTag:

	case AddTag:
	default:
	}
}
