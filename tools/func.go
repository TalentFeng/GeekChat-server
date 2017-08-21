package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

func Md5sum(str string) (string) {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func Password(str string) string  {
	salt := beego.AppConfig.String("salt")
	return Md5sum(str + salt)
}