package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
	_"github.com/go-sql-driver/mysql"
)


type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11; not null; unique)"`
	Password string `gorm:"size:255;not null"`
}

func main(){
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context){
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11{
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"手机号必须为11位"})
			return
		}
		if len(password) < 6{
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码不能少于6位"})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)

		if isTelephoneExist(db, telephone){
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"用户已经存在"})
			return
		}

		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)

		ctx.JSON(200, gin.H{"msg":"注册成功"})

	})
	panic(r.Run())
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}


func RandomString(n int) string {
	var letters = []byte("abfasdfasdfgsadAFASDJGFSLDGJLSDGJNVINR")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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

	db.AutoMigrate(&User{})

	return db
}




























