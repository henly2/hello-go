package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/henly2/hello-go/myproject/models"
)

type RegistController struct {
	beego.Controller
}
type Userinfo struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}
var users Userinfo
//var db *gorm.DB

func (this *RegistController)Regist() {

		data := this.Ctx.Input.RequestBody
		//json数据封装到users对象中
		err1 := json.Unmarshal(data, &users)
		if err1 != nil {
			fmt.Println("json.Unmarshal is err:", err1.Error())
		}
        //把注册信息加入到数据库
		o:=orm.NewOrm()
		user:=models.Userorm{}
		user.Username=users.Username
		user.Password=users.Password
		user.Sex=users.Sex
		user.Age=users.Age
		_, err := o.Insert(&user)
		if err != nil {
			fmt.Println("注册失败")
			return
		}
		this.Ctx.WriteString("注册成功")

}

