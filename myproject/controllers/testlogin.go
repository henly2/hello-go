package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestloginController struct {
	beego.Controller
}
type User1 struct {
	Username string
	Password string
}

var u User1

func (c *TestloginController) Login() {
	name := c.Ctx.GetCookie("name")
	name1 := c.Ctx.Input.Query("name")
	password := c.Ctx.GetCookie("password")
	//c.Ctx.SetCookie("Username",u.Username,-1,"/") 删除cookie
	//c.Ctx.SetCookie("Password",u.Password,-1,"/")
	fmt.Println(name, "-", name1)
	if name != "" {
		c.Ctx.WriteString("Username" + name + "Password" + password)
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_login"method="post">
	                            <input type="text" name="Username"/>
	                            <input type="password" name="Password"/>
	                            <input type="submit" value="提交"/></form></html>`)

	}
}

//id:=c.GetString("id")
//c.Ctx.WriteString("<html>"+id+"<br/>")
//name:=c.Input().Get("name")
//c.Ctx.WriteString(name+"<html/>")

func (c *TestloginController) Post() {

	if err := c.ParseForm(&u); err != nil {

	}
	//写入cookie中
	c.Ctx.SetCookie("Username", u.Username, 100, "/")
	c.Ctx.SetCookie("Password", u.Password, 100, "/")
	//c.SetSession("Username",u.Username)
	//c.SetSession("Password",u.Password)

	c.Ctx.WriteString("Username" + u.Username + "Password" + u.Password)
}
