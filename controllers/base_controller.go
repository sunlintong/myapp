package controllers

import (
	"encoding/hex"
	"crypto/md5"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 密码加密函数，将密码加盐再MD5
func (bc *BaseController) Encode(password string) string {
	salt := "sunlintonghenshuai"
	h := md5.New()
	h.Write([]byte(password+salt))
	return hex.EncodeToString(h.Sum(nil))
}