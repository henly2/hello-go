package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/login1", login1)
	http.HandleFunc("/login2", login2)
	http.ListenAndServe("0.0.0.0:8080", nil)

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("body=", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)

}

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
}

//func mainpost()  {
//	resp,err:=http.Post("127.0.0.1:8000/user/login")
//	if err!=nil{
//		fmt.Println("err=",err)
//		return
//	}
//	defer resp.Body.Close()
//	fmt.Println("body=",resp.Body)
//	body, err := ioutil.ReadAll(resp.Body)
//	fmt.Println(body)
//
//}

//post接口接收json数据
func login1(writer http.ResponseWriter, request *http.Request) {
	var auth Auth
	if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	var result Resp
	if auth.Username == "admin" && auth.Pwd == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "401"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}

//接收x-www-form-urlencoded类型的post请求或者普通get请求
func login2(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username, uError := request.Form["username"]
	pwd, pError := request.Form["password"]

	var result Resp
	if !uError || !pError {
		result.Code = "401"
		result.Msg = "登录失败"
	} else if username[0] == "admin" && pwd[0] == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "203"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

/// 2 ...
//
//// 3 write out
//o := out{
//Id: 1,
//}
//
//od, err := json.Marshal(o)
//if err != nil {
//return
//}
//
//w.Write(od)
//}else {
//return
//}
