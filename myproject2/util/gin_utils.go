package util

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

type UserInfo struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}

var users UserInfo

const (
	SecretKey = "welcome to wangshubo's blog"
)

type User1 struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" bdinding:"required"`
	//Age      int    `form:"age" json:"age"`
}

var user1 User1

type ChangePwd struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

var chpwd ChangePwd

func RegistInfo(c *gin.Context) {
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
func LoginInfo(c *gin.Context) {
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
func Changepwd(c *gin.Context) {
	res_info := c.Request.Header.Get("token")
	_, err := jwt.Parse(res_info, func(token *jwt.Token) (interface{}, error) {

		c.JSON(http.StatusOK, token)
		//claims, ok := token.Claims.(*jwt.MapClaims)
		//if !ok {
		//	return "", fmt.Errorf("error data")
		//}
		//userId, exist := (*claims)["userid"]
		//if !exist {
		//	return "", fmt.Errorf("error data, not has userid")
		//}
		//
		//id, ok := userId.(int)
		//if !ok {
		//	return "", fmt.Errorf("error data, id is not int64")
		//}
		////var user
		//fmt.Println("userId=", id)

		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		err1 := json.Unmarshal(buf[0:n], &chpwd)
		if err1 != nil {
			fmt.Println("json.Unmarshal is err:", err1.Error())
		}
		o := orm.NewOrm()
		user := models.Userorm{}
		//user.Id = id
		user.Username = chpwd.Name
		err2 := o.Read(&user, "username")
		if err2 == nil {
			if user.Password == chpwd.OldPassword {
				user.Password = chpwd.NewPassword
				_, err2 = o.Update(&user)
				if err2 != nil {
					c.JSON(http.StatusOK, "更新失败")
					return "", err2
				}
				c.JSON(http.StatusOK, "更新成功")
			} else {
				c.JSON(http.StatusOK, "更新失败")
			}
		}
		return []byte("SecretKey"), nil
	})

	if err != nil {
		return
	}
	return

}
