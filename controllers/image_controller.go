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
	}else {
		l.Log = "get all images succeed"
	}
	_ = db.InsertLog(l)

	var ret [][4]string
	for _, image := range images {
		for _, rt := range image.RepoTags {
			var data [4]string
			data[0] += rt
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