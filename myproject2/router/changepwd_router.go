package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/models"
	"net/http"
)

type ChangePwd struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

var chpwd ChangePwd

func main() {
	router := gin.Default()
	router.POST("/changepwd", changepwd)
	router.Run(":8080")
}
func changepwd(c *gin.Context) {
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
