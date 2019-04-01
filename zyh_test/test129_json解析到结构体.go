package main

import (
	"encoding/json"
	"fmt"
)

type It1 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"bool"`
}

func main() {
	jsonbuf := `{"company":"itc" ,"subject":{"c++","java"} ,"isok":true}`
	//定义一个结构体变量
	var tmp It1
	err := json.Unmarshal([]byte(jsonbuf), &tmp)
	if err != nil {
		fmt.Println("erro=", err)
		return
	}
}
