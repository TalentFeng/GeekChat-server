package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/astaxie/beego"
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
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Phone string `gorm:"not null;size:11;unique;" valid:"Phone"`
	Mail string `gorm:"size:20;unique;" valid:"Email"`
	Password string `gorm:"not null;size:50 valid:MaxSize(50)"`
	Modified_time time.Time `gorm:"default:current_timestamp"`
}

type UserInfo struct {
	user User
	Profile string `gorm:"not null;size:11;unique"`
	Portrait string `gorm:"size:11;unique"`
	Password string `gorm:"not null;size:50"`
	Modified_time time.Time `gorm:"default:current_timestamp"`
}

func AutoMigrate() {
    db := GetDb()
    db.AutoMigrate(&User{}, &UserInfo{})
}