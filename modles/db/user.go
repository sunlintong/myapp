package db

import (
	"github.com/astaxie/beego/orm"
)

// 用户的数据库表
type User struct {
	ID           int    `orm"column(id);pk;auto"`
	Name         string `orm:"column(name)"`
	Password     string `orm:"column(password)"`
	UserType     string `orm:"column(userType)"`
	RegisterTime int64  `orm:"column(register_time)"`
}

// 创建表User
func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

// 新建用户
func CreateUser(user *User) error {
	o := GetOrmer()
	_, err := o.Insert(user)
	return err
}

// 修改用户信息
func UpdateUser(user *User) error {
	o := GetOrmer()
	_, err := o.Update(user)
	return err
}

// 删除用户
func DeleteUser(user *User) error {
	o := GetOrmer()
	_, err := o.Delete(user)
	return err
}

// 由用户名找出用户
func GetUserByName(name string) (*User, error) {
	user := new(User)
	o := GetOrmer()
	qs := o.QueryTable(new(User))
	err := qs.Filter("name", name).One(user)
	return user, err
}
