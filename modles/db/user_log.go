package db

import (
	"github.com/astaxie/beego/orm"
)

type Log struct {
	ID           int    `orm:"column(id);pk;auto"`
	Time 		 int64  `orm:"column(time)"`
	Name         string `orm:"column(name)"`
	Log			 string `orm:"column(log)"`
}


func InsertLog(l *Log) error {
	o := GetOrmer()
	_, err := o.Insert(l)
	return err
}

// 获取所有log
func GetAllLogs() ([]*Log, error) {
	var logs []*Log
	o := GetOrmer()
	_, err := o.QueryTable(new(Log)).OrderBy("time").All(&logs)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	return logs, err
}