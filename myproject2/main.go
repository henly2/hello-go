package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/util"
)

func main() {
	router := gin.Default()
	router.POST("/regist", util.RegistInfo)
	router.POST("/login", util.LoginInfo)
	router.POST("/changepwd", util.Changepwd)

	router.Run(":8080")
}
