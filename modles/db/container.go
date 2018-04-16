package db

import (
	"log"
	"myapp/modles/local"
	"myapp/types"
)

type Container struct {
	ID           int    `orm:"column(id);pk;auto"`
	Group        string `orm:"column(group)"`
	Container_ID string `orm:"column(container_id)"`
}

const PublicContainerGroup = "public"

func init() {
	o := GetOrmer()
	// 清空
	var is []*Container
	o.QueryTable(new(Container)).All(&is)
	for _, i := range is {
		 o.Delete(i)
	}

	containers, err := local.GetAllContainers()
	if err != nil {
		log.Println(err)
	}
	for _, container := range containers {
		c := new(Container)
		c.Container_ID = container.ID
		c.Group = PublicContainerGroup
		o.Insert(c)
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
