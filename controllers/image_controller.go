package controllers

import (
	"myapp/modles/local"
	"myapp/modles/db"
	"time"
	"fmt"
)

type ImageController struct {
	BaseController
}

func (ic *ImageController) GetAllImages() {
	l := new(db.Log)
	l.Name = "Unknown"
	l.Time = time.Now().Unix()
	images, err := local.GetImages()
	if err != nil {
		l.Log = fmt.Sprintf("get all images failed, %v", err)
	}else {
		l.Log = "get all images succeed"
	}
	_ = db.InsertLog(l)

	var ret [][5]string
	for _, image := range images {
		var data [5]string
		data[0] = fmt.Sprint(image.VirtualSize)
		for _, rt := range image.RepoTags {
			data[1] += rt
		}
		data[2] = image.ID
		data[3] = ic.GetTimeString(image.Created)
		data[4] = fmt.Sprint(image.Size)
		ret = append(ret, data)
	}

	ic.Data["json"] = ret
	ic.ServeJSON()
	ic.Finish()
}