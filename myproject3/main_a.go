package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject3/controller_a"
	log4 "github.com/jeanphorn/log4go"
)

func main() {
	log4.LoadConfiguration("./logininfo.json")
	router := gin.Default()
	router.POST("/user/userregist", controller_a.UserRegist)   //普通用户注册
	router.POST("/user/adminregist", controller_a.AdminRegist) //管理员注册
	router.POST("/user/login", controller_a.UserLogin)         //普通用户登录
	router.POST("/admin/listusers", controller_a.QueryUser)
	router.POST("/admin/loguserlogin", controller_a.LogUserLogin)
	router.Run(":8080")
}
