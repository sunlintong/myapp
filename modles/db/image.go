package db

import (
	"log"
	"myapp/modles/local"
	"myapp/types"
	"time"
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
	for i, image := range images {
		ids[i] = image.Image_ID
	}
	log.Println("镜像ids:", len(ids), ids)
	return ids, err
}

func InsertImage(i *Image) error {
	o := GetOrmer()
	_, err := o.Insert(i)
	return err
}


// 同步数据库与本地image
func SyncImages() {
	o := GetOrmer()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		log.Println("同步镜像。。。。。。。")
		images, err := local.GetImages()
		if err != nil {
			log.Println(err)
		}
		// 检测本地是否新增了容器，若有，添至数据库
		for _, image := range images {
			if !o.QueryTable(new(Image)).Filter("image_id", image.ID).Exist() {
				i := &Image{
					Group: PublicImageGroup,
					Image_ID: image.ID,
				}
				InsertImage(i)
				log.Println("插入新镜像:", i)
			}
		}
		// 检测平台容器是否被删除，若是，从数据库中删除
		u := types.User{
			Name: "admin",
			IsAdmin: true,
		}
		ids, err := GetImageIdsByUser(u)
		log.Println(err)
		for ii, id := range ids {
			var num int64
			for _, i := range images {
				if i.ID == id {
					goto Here
				}
			}
			// 运行到这里，说明找到了待删除的id
			num, _ = o.Delete(&Image{Image_ID: id})
			log.Printf("image 第 %d 行", num)

			Here: 
			log.Println("外循环：", ii)
		}
	}
}
