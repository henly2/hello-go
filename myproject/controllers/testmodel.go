package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/henly2/hello-go/myproject/models"
)

type TestmodelController struct {
	beego.Controller
}

func (c *TestmodelController)Get()  {
	o:=orm.NewOrm()
	user:=models.Userorm{}
	user.Username="s"
	user.Password="s"
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	c.Ctx.WriteString("THIS IS A METHOD Get")
}
