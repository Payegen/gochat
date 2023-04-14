package sql

import (
	"gochat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	user := models.UserBasic{Name: "lxj", PassWord: "123qqq"}
	dsn := "root:123456@tcp(192.168.186.129:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db.AutoMigrate(&models.UserBasic{})
	db.Create(&user)
	return db
}
