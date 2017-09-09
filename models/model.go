package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDb() *gorm.DB {
	conn, _ := beego.AppConfig.GetSection("postgresql")
	var dsn string
	for k, v := range conn {
		dsn += " " + k + "=" + v
	}
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		beego.Error(err)
	}
	is_dev := beego.AppConfig.String("runmode") == "dev"
	db.LogMode(is_dev)
	return db
}

type User struct {
	Id            int       `gorm:"primary_key;AUTO_INCREMENT"`
	Phone         string    `gorm:"not null;size:11;unique;" valid:"Phone" form:"phone"`
	Mail          string    `gorm:"size:20;unique;" valid:"Email" form:"mail"`
	Password      string    `gorm:"not null;size:50 valid:MaxSize(50)" form:"password"`
	Modified_time time.Time `gorm:"default:current_timestamp"`
}

type UserInfo struct {
<<<<<<< HEAD
	Uid           int       `gorm:"primary_key;"`
	NickName      string    `gorm:"size:11;"`
	Profile       string    `gorm:"not null;size:11;unique"`
	Avatar        string    `gorm:"size:100;unique"`
	Sex           int       `gorm:"size:1"`
	Country       int       `gorm:":size:1"`
	Area          int       `gorm:":size:1"`
=======
	Uid           int       `gorm:"primary_key;" form:"uid"`
	Profile       string    `gorm:"not null;size:11;unique" form:"profile"`
	Avatar        string    `gorm:"size:100;unique" form:"avatar"`
>>>>>>> 8f76d5a037ab4586507233e679c173f895c38575
	Modified_time time.Time `gorm:"default:current_timestamp"`
}

type Json struct {
	Code int `json:"code"`
	Msg  string
	Data interface{}
}

func AutoMigrate() {
	db := GetDb()
	db.AutoMigrate(&User{}, &UserInfo{})
	db.Close()
}
