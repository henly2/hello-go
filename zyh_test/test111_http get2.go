package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, error1 := http.Get("https://baidu.com/")
	if error1 != nil {
		fmt.Println("error1=", error1)
	}
	fmt.Println(res)
	body, error2 := ioutil.ReadAll(res.Body)
	if error2 != nil {
		fmt.Println("error2=", error2)
	}
	fmt.Println(string(body))

}
