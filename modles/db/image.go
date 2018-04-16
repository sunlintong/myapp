package db

import (
	"log"
	"myapp/modles/local"
	"myapp/types"
)

type Image struct {
	ID        int    `orm:"column(id);pk;auto"`
	Group string `orm:"column(group)"` //所属用户名
	Image_ID  string `orm:"column(image_id)"`
}

const PublicImageGroup = "public"

func init() {
	o := GetOrmer()
	o.Raw("DROP TABLE `image`").Exec()
	images, err := local.GetImages()
	if err != nil {
		log.Println(err)
	}
	for _, image := range images {
		i := new(Image)
		i.Image_ID = image.ID
		i.Group = PublicImageGroup
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
