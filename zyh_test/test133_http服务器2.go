package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//w给客户回复数据
//r 读取客户端的数据
func HandConn1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	fmt.Println("r.url", r.URL)
	fmt.Println("r.Method=", r.Method)
	fmt.Println("r.Header=", r.Header)
	fmt.Println("r.body=", r.Body)
}
func HandTest(w http.ResponseWriter, r *http.Request) {
	type (
		out struct {
			Id int64 `json:"id"`
		}
	)

	// 3 write out
	o := out{
		Id: 1,
	}

	od, err := json.Marshal(o)
	if err != nil {
		return
	}

	w.Write(od)
}
func HandLogin(w http.ResponseWriter, r *http.Request) {
	type (
		in struct {
			Name string `json:"name"`
		}
		out struct {
			Id int64 `json:"id"`
		}
	)

	// 1 get in
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	i := in{}
	err = json.Unmarshal(d, &i)
	if err != nil {
		return
	}

	// 2 ...

	// 3 write out
	o := out{
		Id: 1,
	}

	od, err := json.Marshal(o)
	if err != nil {
		return
	}

	w.Write(od)
}
func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	//pattern后具体写访问的ip
	http.HandleFunc("/", HandConn1)

	http.HandleFunc("/user/test", HandTest)
	http.HandleFunc("/user/login", HandLogin)
	//监听绑定
	http.ListenAndServe(":8000", nil)

}
