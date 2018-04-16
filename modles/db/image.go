package db

import (
	"myapp/types"
	"myapp/modles/local"
	"log"
)

type Image struct {
	ID         int    `orm:"column(id);pk;auto"`
	Name       string `orm:"column(name)"`	//所属用户名
	Image_ID   string `orm:"column(image_id)"`
	Image_Name string `orm:"column(image_name)"`
}

const PublicImageGroup = "public"

// 由用户获取image_id数组
func GetImageIdsByUser(user types.User) ([]string, error){
	var images []*Image
	var err error
	o := GetOrmer()
	// 管理员看所有
	// 普通用户看public和自己的
	if user.IsAdmin {
		_, err = o.QueryTable(new(Image)).All(&images)
	}else {
		_, err = o.QueryTable(new(Image)).Filter("name__in", user.Name, "public").All(&images)
	}
	ids := make([]string, len(images))
	for _, image := range images {
		id := image.Image_ID
		ids = append(ids, id)
	}
	return ids, err
}

func InsertImage(image *Image) error {
	o := GetOrmer()
	_, err := o.Insert(image)
	return err
}

// 初始化，插入已有镜像,名为public
func InitImageTable() {
	images, err := local.GetImages()
	if err != nil {
		log.Println(err)
	}
	for _, image := range images {
		for _, rt := range image.RepoTags {
			i := new(Image)
			i.Image_Name = rt
			i.Image_ID = image.ID
			i.Name = PublicImageGroup
			InsertImage(i)
		}
	}

}

