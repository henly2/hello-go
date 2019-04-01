package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	js := `{"Age":18,"name":"xiaoming","Addr":{"Province":"Hunan","City":"ChangSha"}}` //name是小写
	//创建一个map
	m := make(map[string]interface{}, 3)
	err := json.Unmarshal([]byte(js), &m) //第二个参数需要地址传递
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("m=", m)

	for key, value := range m {
		//fmt.Println(key,value)
		switch data := value.(type) {

		case string:
			fmt.Println(key, data)

		case float64:
			fmt.Println(key, data)

		case interface{}:
			fmt.Println(key, value)

		}
	}

}
