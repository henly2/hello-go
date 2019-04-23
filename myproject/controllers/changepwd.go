package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	//"github.com/myproject/model_r"
	"fmt"
	"github.com/myproject/models"
)

type Chpwd struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}
type ChangepwdController struct {
	beego.Controller
}

var chpwd Chpwd

func (this *ChangepwdController) Changepwd() (t *jwt.Token) {

	res_info := this.Ctx.Request.Header.Get("token")
	token, err := jwt.Parse(res_info, func(token *jwt.Token) (interface{}, error) {

		this.Ctx.WriteString("token解密成功")
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return "", fmt.Errorf("error data")
		}
		userId, exist := (*claims)["userid"]
		if !exist {
			return "", fmt.Errorf("error data, not has userid")
		}

		id, ok := userId.(int)
		if !ok {
			return "", fmt.Errorf("error data, id is not int64")
		}
		//var user
		fmt.Println("userId=", id)

		data := this.Ctx.Input.RequestBody
		//	//json数据封装到user对象中
		err := json.Unmarshal(data, &chpwd)
		if err != nil {
			return "", err
		}
		o := orm.NewOrm()
		user := models.Userorm{}
		user.Id = id
		user.Username = chpwd.Name
		err2 := o.Read(&user, "username")
		if err2 == nil {
			if user.Password == chpwd.OldPassword {
				user.Password = chpwd.NewPassword
				_, err2 = o.Update(&user)
				if err2 != nil {
					this.Ctx.WriteString("更新失败")
					return "", err2
				}
				this.Ctx.WriteString("更新成功")
			} else {
				this.Ctx.WriteString("更新失败")
			}
		}
		return []byte("SecretKey"), nil
	})

	if err != nil {
		return token
	}
	return

}
