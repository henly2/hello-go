package controllers

import (
	"github.com/astaxie/beego"
)

type TestinputController struct {
	beego.Controller
}
type User struct {
	Username string
	Password string
}

func (c *TestinputController) Get() {

	//id:=c.GetString("id")
	//c.Ctx.WriteString("<html>"+id+"<br/>")
	//name:=c.Input().Get("name")
	//c.Ctx.WriteString(name+"<html/>")

	//name:=c.GetSession("Username")
	//password:=c.GetSession("Password")
	//if namestring,ok:=name.(string);ok &&namestring!=""{
	//	//c.Ctx.WriteString("Username"+name+"Password"+password)
	//	c.Ctx.WriteString("name:"+name.(string)+"password:"+password.(string))
	//}else{
	c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input"method="post">
                                <input type="text" name="username"/>
                                <input type="password" name="password"/>
                                <input type="submit" value="提交"/></form></html>`)
	//}

}
func (c *TestinputController) Post() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {

	}
	c.Ctx.WriteString("Username:" + u.Username + "Password:" + u.Password)
}
