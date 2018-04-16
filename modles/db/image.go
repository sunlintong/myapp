package db

import (
	"log"
	"myapp/modles/local"
	"myapp/types"
)

type Image struct {
	ID       int    `orm:"column(id);pk;auto"`
	Group    string `orm:"column(group)"`
	Image_ID string `orm:"column(image_id)"`
}

const PublicImageGroup = "public"

// 初始化Image表
func init() {
	o := GetOrmer()
	// 先清空
	var is []*Image
	o.QueryTable(new(Image)).All(&is)
	for _, i := range is {
		 o.Delete(i)
	}
	images, err := local.GetImages()
	if err != nil {
		log.Println(err)
	}
	for _, image := range images {
		i := new(Image)
		i.Image_ID = image.ID
		i.Group = PublicImageGroup
		o.Insert(i)
	}
}

// 由用户获取image_id数组
func GetImageIdsByUser(user types.User) ([]string, error) {
	var images []*Image
	var err error
	o := GetOrmer()
	// 管理员看所有
	// 普通用户看public和自己的
	if user.IsAdmin {
		_, err = o.QueryTable(new(Image)).All(&images)
	} else {
		_, err = o.QueryTable(new(Image)).Filter("group__in", user.Name, "public").All(&images)
	}
	ids := make([]string, len(images))
	for _, image := range images {
		for i := range ids {
			ids[i] = image.Image_ID
		}
	}
	log.Println("镜像ids:", len(ids), ids)
	return ids, err
}
