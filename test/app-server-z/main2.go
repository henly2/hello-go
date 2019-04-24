package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
)

//定义模型
type User struct {
	gorm.Model
	//Id int
	Username string `gorm:"column:username"`
	Age      int    `gorm:"column:age"`
	Sex      string `gorm:"column:sex"`
	Password string `gorm:"column:password"`
}

var db *gorm.DB

/*
目标：实现一个用户系统
用户接口：注册，登陆，修改密码
初始管理员接口：登陆，修改密码，创建子管理员，删除子管理员
子管理员接口：登陆，修改密码，查看用户列表

初始管理员有默认的初始账户和密码

接口：http
配置文件：yaml或者json
数据库：mysql
缓存：redis

框架：
http框架： beego, iris， go-gin 等
数据库框架：orm
缓存: go-redis
*/

//注册
func Regist(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" { //判断是不是POST请求
		userinfo, _ := ioutil.ReadAll(req.Body) //读取数据
		HandRegist(userinfo)
		res.WriteHeader(200)
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//用户注册
func HandRegist(userInfo []byte) {
	type Auth struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Password string `json:"password"`
	}
	var user1 Auth
	err := json.Unmarshal(userInfo, &user1) //将 json 转换成结构体
	if err != nil {
		fmt.Println("err=", err, "-", string(userInfo))
		return
	}

	//添加数据到数据库
	users := &User{Username: user1.Username, Age: user1.Age, Sex: user1.Sex, Password: user1.Password}
	db.Create(users) //插入数据到数据库
	return
}

//登录
func Login(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" { //判断是不是POST请求
		HandLogin(req) //登录
		res.WriteHeader(200)
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//登录
func HandLogin(req *http.Request) {
	type in struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	inf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	i := in{}
	err = json.Unmarshal(inf, &i)
	if err != nil {
		return
	}
	fmt.Println("i.name=", i.Name, "password=", i.Password)

	//从数据库查询数据，
	//查询所有的数据
	var user []User

	db.Model(&User{}).Where("username=?", i.Name).Find(&user)

	for _, a := range user {
		if a.Username == i.Name {
			if a.Password == i.Password {
				fmt.Println("登录成功")
				return
			} else {
				fmt.Println("登录失败")
				return
			}
		}
	}

	return
}

//修改密码
func ChangePwd(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" { //判断是不是POST请求
		ChangeInfo(req) //登录
		res.WriteHeader(200)
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//修改密码
func ChangeInfo(req *http.Request) (info interface{}) {
	type in struct {
		Name        string `json:"name"`
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}
	inf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	i := in{}
	err = json.Unmarshal(inf, &i)
	if err != nil {
		return
	}
	fmt.Println("i.name=", i.Name, "i.Oldpassword=", i.OldPassword)

	//修改密码
	//判断旧密码是否正确
	var users []User
	db.Model(&User{}).Where("username=?", i.Name).Find(&users)
	for _, a := range users {
		if a.Username == i.Name {
			if a.Password == i.OldPassword {
				db.Model(&User{}).Update("username", i.NewPassword) //修改旧密码
				fmt.Println("修改成功")
				return
			} else {
				fmt.Println("修改失败")
				return
			}
		}
	}
	return
}

//返回信息
func Returninfo(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("点击此接口返回信息"))
	fmt.Println(req.Header)
}

func main() {
	//连接数据库
	newdb, err := gorm.Open("mysql", "root:mysql@/golang_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
		newdb.CreateTable(&User{}) //创建表，默认表名是users
	}
	db = newdb

	//用户注册
	http.HandleFunc("/user/regist", Regist)

	//用户登录
	http.HandleFunc("/user/login", Login)

	//用户修改密码
	http.HandleFunc("/user/changePwd", ChangePwd)

	//返回信息
	http.HandleFunc("/user/getdata", Returninfo)

	//监听绑定
	http.ListenAndServe(":8000", nil)

}
