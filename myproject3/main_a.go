package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject3/controller_a"
	log4 "github.com/jeanphorn/log4go"
)

func main() {
	log4.LoadConfiguration("./logininfo.json")
	router := gin.Default()
	v1 := router.Group("admin")
	{
		v1.POST("/regist", controller_a.AdminRegist)
		v1.POST("/login", controller_a.AdminLogin)
		v1.POST("/loguserlogin", controller_a.LogAdminLogin)
	}
	v2 := router.Group("user")
	{
		v2.POST("/regist", controller_a.UserRegist)
		v2.POST("/login", controller_a.UserLogin)
		v2.POST("/listusers", controller_a.QueryAdmin)
	}

	router.Run(":8080")
}
