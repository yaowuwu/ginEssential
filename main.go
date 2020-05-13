package main

import (
	"github.com/gin-gonic/gin"
	"hello/ginessential/common"
	_"github.com/go-sql-driver/mysql"
)



func main(){
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run())
}






























