package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company string
	Subject []string
	Isok    bool
	price   float64
}

func main() {
	//定义一个结构体变量同时初始化
	s := It{"it", []string{"c++", "go", "python"}, true, 12.1}
	//编码，生成json文本
	str, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("str=", string(str))

}
