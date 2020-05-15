package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hello/ginessential/common"
	"hello/ginessential/model"
	"hello/ginessential/response"
	"hello/ginessential/vo"
	"strconv"
)

//定义完接口, 再让结构体实现接口

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController{
	db :=  common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB:db}
}


func (c CategoryController) Create(ctx *gin.Context) {
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)
	//if requestCategory.Name == "" {
	//	response.Fail(ctx, nil,"数据验证错误, 分类名称必填")
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx,nil, "数据验证错误,分类名称必填")
		return
	}

	category := model.Category{Name:requestCategory.Name}
	c.DB.Create(&category)
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)
	//if requestCategory.Name == ""{
	//	response.Fail(ctx,nil,"数据验证错误, 分类名必填")
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx,nil, "数据验证错误,分类名称必填")
		return
	}


	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound(){
		response.Fail(ctx,nil, "分类不存在")
		return
	}
	//更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category":updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var Category model.Category
	if c.DB.First(&Category, categoryId).RecordNotFound(){
		response.Fail(ctx,  nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category":Category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId , _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil{
		response.Fail(ctx, nil, "删除失败, 请重试")
		return
	}
	response.Success(ctx,nil,"删除成功")
}


