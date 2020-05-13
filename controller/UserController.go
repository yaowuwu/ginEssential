package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"hello/ginessential/common"
	"hello/ginessential/dto"
	"hello/ginessential/model"
	"hello/ginessential/response"
	"hello/ginessential/util"
	"log"
	"net/http"
)

func Register(ctx *gin.Context){
	DB := common.GetDB()

	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"手机号必须为11位"})
		return
	}
	if len(password) < 6{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码不能少于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	if isTelephoneExist(DB, telephone){
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"用户已经存在"})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code":500, "msg":"加密错误"})
		return
	}

	newUser := model.User{
		Name: name,
		Telephone: telephone,
		//Password: hasedPassword,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	response.Success(ctx,nil, "注册成功")
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":"注册成功"})

}


func Login(ctx *gin.Context){
	DB := common.GetDB()

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

	var user model.User
	DB.Where("telephone = ?" , telephone).First(&user)
	if user.ID == 0{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"用户不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"code":400, "msg":"密码错误"})
		return
	}

	token,err := common.ReleaseToken(user)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"code":500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	response.Success(ctx,gin.H{"token":token} ,"登录成功")
	//ctx.JSON(200, gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	//	"msg": "登录成功",
	//})

}

//var a interface{} = int(10)
//var b MyInt = a.(MyInt)

func Info(ctx *gin.Context){
	user, _ := ctx.Get("user")

	//cannot use user (type interface {}) as type model.User in argument to dto.ToUserDto: need type assertion
	ctx.JSON(http.StatusOK, gin.H{"code":200, "data":gin.H{"user": dto.ToUserDto(user.(model.User))}})
}


func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}