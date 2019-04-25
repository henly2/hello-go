package util

import (
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

const (
	SecretKey = "hello world"
)

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" bdinding:"required"`
}
type Changepwd struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

//用户注册
func RegistInfo(c *gin.Context) {

	var users UserInfo
	err := c.BindJSON(&users)
	if err != nil {
		c.JSON(http.StatusOK, "错误数据")
		return
	}
	//把注册信息加入到数据库
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = users.Username
	user.Password = users.Password
	user.Sex = users.Sex
	user.Age = users.Age
	_, err = o.Insert(&user)
	if err != nil {
		c.JSON(http.StatusOK, "注册失败")
		return

	}
	c.JSON(http.StatusOK, "注册成功")

}

//登录
func LoginInfo(c *gin.Context) {

	var users User
	err := c.BindJSON(&users)
	if err != nil {
		c.JSON(http.StatusOK, "错误数据")
		return
	}
	//从数据库查询数据，
	//查询所有的数据
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = users.Name
	err = o.Read(&user, "username")
	if err != nil {
		c.JSON(http.StatusOK, "登陆失败")
		return
	} else {
		if user.Username == users.Name && user.Password == users.Password {

			//登陆成功生成token
			claims := make(jwt.MapClaims)
			claims["userid"] = user.Id
			claims["exp"] = time.Now().Add(time.Hour * 48).Unix() //设置过期时间，过期需要重新获取
			claims["iat"] = time.Now().Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(SecretKey)) //使用自定义字符串进行加密
			if err != nil {
				c.JSON(http.StatusOK, "加密失败")
				return
			}

			c.Header("token", tokenString)
			c.JSON(http.StatusOK, "登录成功")
			return
		} else {
			c.JSON(http.StatusOK, "登录失败")
			return
		}
	}
}

//修改密码
func ChangePwd(c *gin.Context) {
	var chpwd Changepwd

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

		err := c.BindJSON(&chpwd)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return "", err
		}
		o := orm.NewOrm()
		user := models.Userorm{}
		//user.Id = id

		user.Username = chpwd.Name
		err = o.Read(&user, "username")
		if err == nil {
			if user.Password == chpwd.OldPassword {
				user.Password = chpwd.NewPassword
				_, err = o.Update(&user)
				if err != nil {
					c.JSON(http.StatusOK, "更新失败")
					return "", err
				}
				c.JSON(http.StatusOK, "更新成功")
			} else {
				c.JSON(http.StatusOK, "更新失败")
			}
		}
		return []byte("SecretKey"), nil
	})
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

}
