package controllers

import (
	"encoding/json"
	"fmt"
	"myapp/modles/db"
	"myapp/modles/local"
	"time"
	"log"
	//"github.com/docker/docker/api/types"
)

// 时间类型
const (
	Remove = "remove"
	ForceRemove = "force"
	RemoveAll = "removeall"
	ForceAll = "forceall"
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
	l.Name = ic.User.Name
	l.Time = time.Now().Unix()
	ids, err := db.GetImageIdsByUser(ic.User)
	log.Println("CONTROLLER IMAGE IDS: ", len(ids), (ids))
	ic.CheckErr(err)
	images, err := local.GetImagesByImageIds(ids)
	log.Println("CONTROLLER IMAGES:", len(images))
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
			if rt == "<none>:<none>" {
				data[0] = "未命名"
			} else {
				data[0] = rt
			}
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
	l.Name = ic.User.Name
	l.Time = time.Now().Unix()

	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &req)
	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ic.BadRequest(l)
		return
	}

	var resp [][3]string
	var str [3]string
	// 事件的逻辑处理部分
	switch req.Event_Type {
	case Remove:
		items, err := local.RemoveImage(req.Image_ID)
		if err != nil {
			l.Log = err.Error()
			db.InsertLog(l)
			ic.ServiceError(l)
			return
		}
		for _, item := range items {
			str[0] = "DELETE："
			str[1] = item.Deleted
			str[2] = item.Untagged
			resp = append(resp, str)
		}

	case ForceRemove:
		items, err := local.ForceRemoveImage(req.Image_ID)
		if err != nil {
			l.Log = err.Error()
			db.InsertLog(l)
			ic.ServiceError(l)
			return
		}
		for _, item := range items {
			str[0] = "DELETE："
			str[1] = item.Deleted
			str[2] = item.Untagged
			resp = append(resp, str)
		}

	case RemoveAll:
		items, err := local.RemoveImageAndChildren(req.Image_ID)
		if err != nil {
			l.Log = err.Error()
			db.InsertLog(l)
			ic.ServiceError(l)
			return
		}
		for _, item := range items {
			str[0] = "DELETE："
			str[1] = item.Deleted
			str[2] = item.Untagged
			resp = append(resp, str)
		}

	case ForceAll:
		items, err := local.ForceRemoveImageAndChildren(req.Image_ID)
		if err != nil {
			l.Log = err.Error()
			db.InsertLog(l)
			ic.ServiceError(l)
			return
		}
		for _, item := range items {
			str[0] = "DELETE："
			str[1] = item.Deleted
			str[2] = item.Untagged
			resp = append(resp, str)
		}

	default:
		err = fmt.Errorf("unknown event %v", req.Event_Type)
	}

	if err != nil {
		l.Log = err.Error()
		db.InsertLog(l)
		ic.ServiceError(l)
	} else {
		l.Log = fmt.Sprintf("remove image success, request:%+v", req)
		db.InsertLog(l)
		ic.Success(resp)
	}

}