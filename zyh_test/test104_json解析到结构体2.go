package main

import (
	"encoding/json"
	"fmt"
)

type addr struct {
	Province string
	City     string
}
type stu struct {
	Name string
	Age  int
	Addr addr
}

func main() {
	js := `{"Age":18,"name":"xiaoming","Addr":{"Province":"Hunan","City":"ChangSha"}}` //name是小写
	var xm stu
	err := json.Unmarshal([]byte(js), &xm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xm) //输出{xiaoming 18 {Hunan ChangSha}}
}
