package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hello/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB{
	driveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db , err := gorm.Open(driveName, args)
	if err != nil{
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB{
	return DB
}