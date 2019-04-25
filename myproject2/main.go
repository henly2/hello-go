package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/controller"
)

func main() {
	router := gin.Default()
	router.POST("/regist", controller.UserRegist)
	router.POST("/login", controller.UserLogin)
	router.POST("/changepwd", controller.UserChangePassword)
	router.Run(":8080")
}
