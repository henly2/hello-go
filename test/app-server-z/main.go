package main

//
//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	_ "github.com/go-sql-driver/mysql"
//)
//var (
//	dbhostsip  = "127.0.0.1:3306"//IP地址
//	dbusername = "root@localhost"//用户名
//	dbpassword = "mysql"//密码
//	dbname     = "students"//表名
//)
//
//type mysql_db struct {
//	db *sql.DB
//}
//
///*
//目标：实现一个用户系统
//用户接口：注册，登陆，修改密码
//初始管理员接口：登陆，修改密码，创建子管理员，删除子管理员
//子管理员接口：登陆，修改密码，查看用户列表
//
//初始管理员有默认的初始账户和密码
//
//接口：http
//配置文件：yaml或者json
//数据库：mysql
//缓存：redis
//
//框架：
//http框架： beego, iris， go-gin 等
//数据库框架：orm
//缓存: go-redis
//*/
//
////var p map[string]interface{}
//var f mysql_db
//
////注册
//func Regist(res http.ResponseWriter, req *http.Request) {
//
//	if req.Method == "POST" {  //判断是不是POST请求
//		userinfo, _ := ioutil.ReadAll(req.Body) //读取数据
//		f.HandRegist(userinfo)
//		res.WriteHeader(200)
//	} else {
//		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
//	}
//}
//
////用户注册
//func (f *mysql_db) HandRegist(userInfo[]byte) (a map[string]string) {
//	type Auth struct {
//		UserName string `json:"username"`
//		Age int  `json:"age"`
//		Sex string `json:"sex"`
//		Password int `json:"password"`
//	}
//	var user Auth
//	err:=json.Unmarshal(userInfo, &user) //将 json 转换成结构体
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//添加数据到数据库
//	stmt, err := f.db.Prepare("insert into user(username,password) VALUES(?,?)")
//	checkErr(err)
//	res, err := stmt.Exec(user.UserName,user.Password)
//	checkErr(err)
//	id, err := res.LastInsertId()
//	checkErr(err)
//	fmt.Println(id)
//
//	//p[user.UserName]=user   //把用户信息存到map中
//	return
//}
//
////登录
//func Login(res http.ResponseWriter, req *http.Request) {
//
//	if req.Method == "POST" {  //判断是不是POST请求
//		f.HandLogin(req)  //登录
//		res.WriteHeader(200)
//	} else {
//		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
//	}
//}
//
////登录
//func (f *mysql_db) HandLogin(req *http.Request)(info interface{}){
//	type in struct {
//			Name string `json:"name"`
//			Password string `json:"password"`
//		}
//	inf, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		return
//	}
//	i := in{}
//	err = json.Unmarshal(inf, &i)
//	if err != nil {
//		return
//	}
//	fmt.Println("i.name=",i.Name,"i.password=",i.Password)
//
//	//从数据库查询数据，
//	rows,err:=f.db.Query("select username,password from user")
//	for rows.Next(){
//		var userName string
//		var password string
//
//		err=rows.Scan(&userName,&password)
//		checkErr(err)
//		fmt.Println(userName)
//		fmt.Println(password)
//		if i.Name==userName{
//			if i.Password==password{
//				return "登录成功"
//			}
//		} else {
//			return "登录失败"
//		}
//	}
//	return
//}
//
////修改密码
//func ChangePwd(res http.ResponseWriter, req *http.Request) {
//	if req.Method == "POST" {  //判断是不是POST请求
//		f.ChangeInfo(req)  //登录
//		res.WriteHeader(200)
//	} else {
//		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
//	}
//}
//
////修改密码
//func (f *mysql_db) ChangeInfo(req *http.Request)(info interface{}) {
//	type in struct {
//		Name string `json:"name"`
//		OldPassword string `json:"oldpassword"`
//		NewPassword string  `json:"newpassword"`
//	}
//	inf, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		return
//	}
//	i := in{}
//	err = json.Unmarshal(inf, &i)
//	if err != nil {
//		return
//	}
//	fmt.Println("i.name=",i.Name,"i.Oldpassword=",i.OldPassword)
//
//	//查询当用户名是否存在以及密码是否正确
//	rows,err:=f.db.Query("select username,password from user")
//	for rows.Next(){
//		var userName string
//		var password string
//
//		err=rows.Scan(&userName,&password)
//		checkErr(err)
//		if i.Name==userName{
//			if i.OldPassword==password{
//				//更改密码
//				stmt,err:=f.db.Prepare("update user set password where username=?")
//				res,err:=stmt.Exec(i.NewPassword,i.Name)
//				checkErr(err)
//				affect,err:=res.RowsAffected()
//				checkErr(err)
//				fmt.Println(affect)
//
//			}
//		} else {
//			return "修改失败"
//		}
//	}
//	return
//}
//
////返回信息
//func Returninfo (res http.ResponseWriter, req *http.Request) {
//	res.Write([]byte("点击此接口返回信息"))
//	fmt.Println(req.Header)
//}
//func main()  {
//	f.mysql_open()
//	//用户注册
//	http.HandleFunc("/user/regist",Regist)
//
//	//用户登录
//	http.HandleFunc("/user/login",Login)
//
//	//用户修改密码
//	http.HandleFunc("/user/changePwd",ChangePwd)
//
//	//返回信息
//	http.HandleFunc("/user/getdata",Returninfo)
//
//	//监听绑定
//	http.ListenAndServe(":8000",nil)
//
//
//}
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (f *mysql_db) mysql_open() {
//
//	//type User struct {
//	//	gorm.Model
//	//	//Id int
//	//	Username string
//	//	Age int
//	//	Sex string
//	//	Password string
//	//}
//	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/gomysql?charset=utf8")
//	//checkErr1(err)
//	//db, err := gorm.Open("mysql", "root:mysql@/golang_gorm?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		fmt.Println("连接失败",err)
//	}
//	fmt.Println("连接成功",db)
//
//
//
//}
