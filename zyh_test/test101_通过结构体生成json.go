package main

import (
	"encoding/json"
	"fmt"
)

//成员首字母必须大写
type IT struct {
	Company  string
	Subjects []string
	Isok     bool
	price    float64
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"it", []string{"go", "python", "c++", "test"}, true, 22.2}
	//编码，根据内容生成json文本
	str, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("str=", string(str))

}
