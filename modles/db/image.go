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
	var ids []string
	var err error
	o := GetOrmer()
	if user.IsAdmin {
		_, err = o.QueryTable(new(Image)).All(&ids)
	}else {
		_, err = o.QueryTable(new(Image)).Filter("name", user.Name).All(&ids)
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

