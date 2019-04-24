package controllers

import (
	//"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/myproject1/models"
	"net/http"
	"strconv"
	"time"
)

type AppClaims struct {
	UserId uint `json:"uid"`
}

const (
	SecretKey = "welcome to wangshubo's blog"
)

type Token1 struct {
	Token string `json:"token"`
}

type In struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginController struct {
	beego.Controller
}

var user1 In

func (this *LoginController) Login() {

	//从请求中读取数据
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err1 := json.Unmarshal(data, &user1)
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
		this.Ctx.WriteString("登录失败")
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
				this.Ctx.WriteString("加密失败")
				return
			}
			if err != nil {
				this.Ctx.Output.Header("SetStatus", strconv.Itoa(http.StatusInternalServerError))
				//fatal(err)
				http.Error(this.Ctx.ResponseWriter, "Server is Wrong", http.StatusInternalServerError)
				return
			}

			fmt.Println("Token:", tokenString)

			this.Ctx.ResponseWriter.Header().Set("token", tokenString)
			this.Ctx.WriteString("登录成功")
			this.Ctx.WriteString(fmt.Sprintf("{\"Token\":\"%s\"}", tokenString))
			//log.Println("aaaaa")
			return
		} else {
			this.Ctx.WriteString("登录失败")
			return
		}
	}
}
