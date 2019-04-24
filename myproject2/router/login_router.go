package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/models"
	"net/http"
	"time"
)

const (
	SecretKey = "welcome to wangshubo's blog"
)

type User1 struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" bdinding:"required"`
	//Age      int    `form:"age" json:"age"`
}

var user1 User1

func main() {
	router := gin.Default()
	router.POST("/login", logininfo)
	router.Run(":8080")
}
func logininfo(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	err1 := json.Unmarshal(buf[0:n], &user1)
	if err1 != nil {
		fmt.Println("json.Unmarshal is err:", err1.Error())
	}
	//从数据库查询数据，
	//查询所有的数据
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = user1.Name
	err := o.Read(&user, "username")
	if err != nil {
		c.JSON(http.StatusOK, "登陆成功")
		return
	} else {
		if user.Username == user1.Name && user.Password == user1.Password {

			//登陆成功生成token
			claims := make(jwt.MapClaims)
			claims["userid"] = user.Id
			claims["exp"] = time.Now().Add(time.Hour * 48).Unix() //设置过期时间，过期需要重新获取
			claims["iat"] = time.Now().Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//token.Claims = claims

			tokenString, err2 := token.SignedString([]byte(SecretKey)) //使用自定义字符串进行加密

			if err2 != nil {
				c.JSON(http.StatusOK, "加密失败")
				return
			}
			//if err != nil {
			//	c.JSON(http.StatusOK, strconv.Itoa(http.StatusInternalServerError))
			//	//fatal(err)
			//	//http.Error(this.Ctx.ResponseWriter, "Server is Wrong", http.StatusInternalServerError)
			//	return
			//}

			fmt.Println("Token:", tokenString)
			c.Writer.Header().Set("token", tokenString)
			c.JSON(http.StatusOK, "登录成功")
			//this.Ctx.WriteString("登录成功")
			c.JSON(http.StatusOK, tokenString)
			//this.Ctx.WriteString(fmt.Sprintf("{\"Token\":\"%s\"}", tokenString))
			//log.Println("aaaaa")
			return
		} else {
			c.JSON(http.StatusOK, "登录失败")
			return
		}
	}

}
