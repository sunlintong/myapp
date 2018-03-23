package controllers

import (
	"encoding/json"
	"fmt"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"github.com/docker/docker/api/types"
)

// 时间类型
const (
	rmImage = "remove"
)

type ImageRequest struct {
	Image_ID   string `json:"image_id"`
	Event_Type string `json:"event_type"`
}

type ImageController struct {
	BaseController
}

func (ic BaseController) Get() {
	ic.TplName = "image.html"
}

func (ic *ImageController) GetAllImages() {
	l := new(db.Log)
	l.Name = "Unknown"
	l.Time = time.Now().Unix()
	images, err := local.GetImages()
	if err != nil {
		l.Log = fmt.Sprintf("get all images failed, %v", err)
	} else {
		l.Log = "get all images succeed"
	}
	_ = db.InsertLog(l)

	var ret [][4]string
	for _, image := range images {
		for _, rt := range image.RepoTags {
			var data [4]string
			data[0] += rt
			// 剪去前面的 "sha256:" 字符串
			image.ID = string([]byte(image.ID)[7:])
			data[1] = image.ID
			data[2] = ic.GetTimeString(image.Created)
			data[3] = fmt.Sprintf("%.2f", float64(image.Size)/1000000) + " MB"
			ret = append(ret, data)
		}
	}

	ic.Data["json"] = ret
	ic.ServeJSON()
	ic.Finish()
}

// 处理镜像相关操作 POST
func (ic *ImageController) OperateImage() {
	var req ImageRequest
	l := new(db.Log)
	l.Name = "unknown"
	l.Time = time.Now().Unix()

	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ic.BadRequest(l)
		return
	}

	var resp []types.ImageDeleteResponseItem
	// 事件的逻辑处理部分
	switch req.Event_Type {
	case rmImage:
		resp, err = local.RemoveImage(req.Image_ID)
	default:
		err = fmt.Errorf("unknown event %v", req.Event_Type)
	}

	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ic.ServiceError(l)
	} else {
		l.Log = fmt.Sprintf("v%", resp)
		db.InsertLog(l)
		ic.Success(resp)
	}

}
