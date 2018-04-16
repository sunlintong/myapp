package db

type Container struct {
	ID           int    `orm:"column(id);pk;auto"`
	Name         string `orm:"column(name)"`
	Container_ID string `orm:"column(container_id)"`
}
