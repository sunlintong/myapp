package controllers

import (
	"encoding/hex"
	"crypto/md5"
	"github.com/astaxie/beego"
	"time"
	"myapp/types"
)

type BaseController struct {
	beego.Controller
	User types.User
}

// 密码加密函数，将密码加盐再MD5
func (bc *BaseController) Encode(password string) string {
	salt := "sunlintonghenshuai"
	h := md5.New()
	h.Write([]byte(password+salt))
	return hex.EncodeToString(h.Sum(nil))
}

// 将unix时间戳格式化为字符串
func (bc *BaseController) GetTimeString(now int64) string {
	return time.Unix(now, 0).Format("2006-01-02 15:04:05")
}