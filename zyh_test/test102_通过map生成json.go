package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "it"
	m["subject"] = []string{"go", "c++", "python"}
	m["isok"] = true
	m["price"] = 22.1
	//编码成json
	result, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}
