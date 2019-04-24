package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Regist1() {
	type user struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Password string `json:"password"`
	}

	tmp := user{"xx", 5, "男", "xx"}
	data, err := json.Marshal(tmp) //转化为JSON
	if err != nil {
		fmt.Println("json Marshal err", err)
	}
	resp, err := http.Post("http://127.0.0.1:8000/user/regist", "application/json", bytes.NewBuffer([]byte(data)))
	fmt.Println(data)
	fmt.Println(bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //读取服务器返回的信息
	if err != nil {
		fmt.Println("read err")
	}
	fmt.Println(string(body))

}
func Login1() {
	type user1 struct {
		Name string `json:"name"`

		Password string `json:"password"`
	}
	tmp := user1{"xx", "xx"}
	data, err := json.Marshal(tmp) //转化为JSON
	if err != nil {
		fmt.Println("json Marshal err", err)
	}
	resp, err := http.Post("http://127.0.0.1:8000/user/login", "application/json", bytes.NewBuffer([]byte(data)))
	fmt.Println(data)
	fmt.Println(bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //读取服务器返回的信息
	if err != nil {
		fmt.Println("read err")
	}
	fmt.Println(string(body))
}
func ChangePwd1() {
	type user2 struct {
		Name        string `json:"name"`
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}
	tmp := user2{"xx", "xx", "ss"}
	data, err := json.Marshal(tmp) //转化为JSON
	if err != nil {
		fmt.Println("json Marshal err", err)
	}
	resp, err := http.Post("http://127.0.0.1:8000/user/changePwd", "application/json", bytes.NewBuffer([]byte(data)))
	fmt.Println(data)
	fmt.Println(bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //读取服务器返回的信息
	if err != nil {
		fmt.Println("read err")
	}
	fmt.Println(string(body))
}
func main() {
	//Regist1()
	//Login1()
	ChangePwd1()

}
