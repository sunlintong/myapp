package db

import (
	"log"
	"myapp/modles/local"
	"myapp/types"
	"time"
)

type Container struct {
	ID           int    `orm:"column(id);pk;auto"`
	Group        string `orm:"column(group)"`
	Container_ID string `orm:"column(container_id)"`
}

const PublicContainerGroup = "public"

func init() {
	o := GetOrmer()
	var is []*Container
	o.QueryTable(new(Container)).All(&is)
	for _, i := range is {
		 o.Delete(i)
	}
	containers, err := local.GetAllContainers()
	if err != nil {
		log.Println(err)
	}
	// 检测本地是否新增了容器，若有，添至数据库
	for _, container := range containers {
		if !o.QueryTable(new(Container)).Filter("container_id", container.ID).Exist() {
			c := &Container{
				Group: PublicContainerGroup,
				Container_ID: container.ID,
			}
			InsertContainer(c)
		}
	}
}

// 由用户获取image_id数组
func GetContainerIdsByUser(user types.User) ([]string, error) {
	var containers []*Container
	var err error
	o := GetOrmer()
	// 管理员看所有
	// 普通用户看public和自己的
	if user.IsAdmin {
		_, err = o.QueryTable(new(Container)).All(&containers)
	} else {
		_, err = o.QueryTable(new(Container)).Filter("group__in", user.Name, "public").All(&containers)
	}
	log.Println("ccccc", len(containers), containers)
	ids := make([]string, len(containers))
	for i, container := range containers {
		ids[i] = container.Container_ID
	}
	log.Println("cid:", len(ids), ids)
	return ids, err
}

func InsertContainer(c *Container) error {
	o := GetOrmer()
	_, err := o.Insert(c)
	return err
}

// 同步数据库与本地container
func SyncContainers() {
	log.Println("SYNCCONTAINER")
	o := GetOrmer()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		containers, err := local.GetAllContainers()
		if err != nil {
			log.Println(err)
		}
		// 检测本地是否新增了容器，若有，添至数据库
		for _, container := range containers {
			if !o.QueryTable(new(Container)).Filter("container_id", container.ID).Exist() {
				c := &Container{
					Group: PublicContainerGroup,
					Container_ID: container.ID,
				}
				InsertContainer(c)
				log.Println("插入container:", c)
			}
		}
		// 检测平台容器是否被删除，若是，从数据库中删除
		u := types.User{
			Name: "admin",
			IsAdmin: true,
		}
		ids, err := GetContainerIdsByUser(u)
		log.Println(err)
		for _, id := range ids {
			var num int64
			for _, c := range containers {
				if c.ID == id {
					goto Here
				}
			}
			// 运行到这里，说明找到了待删除的id
			num, _ = o.Delete(&Container{Container_ID: id})
			log.Printf("container %s 不见了，删除数据库中container 第 %d 行", id, num)
			Here: 
		}
	}
}