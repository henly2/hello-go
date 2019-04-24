package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/models"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/regist", registinfo)
	router.Run(":8080")
}

type UserInfo struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}

var users UserInfo

func registinfo(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	err1 := json.Unmarshal(buf[0:n], &users)
	if err1 != nil {
		fmt.Println("json.Unmarshal is err:", err1.Error())
	}
	//把注册信息加入到数据库
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = users.Username
	user.Password = users.Password
	user.Sex = users.Sex
	user.Age = users.Age
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("注册失败")
		return
	}
	c.JSON(http.StatusOK, "注册成功")

}
