package main

import (
	"encoding/json"
	"fmt"
)

type IT1 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonbuf := `
{
  "company":"itc",
  "subjects":{
  "go",
  "c++"
    },
    "isok":true,
    "price":66.6
}`
	//定义一个结构体变量
	var tmp IT1
	err := json.Unmarshal([]byte(jsonbuf), &tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmp)

}
