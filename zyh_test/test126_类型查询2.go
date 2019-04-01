package main

import (
	"fmt"
)

type Students6 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "aa"
	i[2] = Students6{"adf", 5}
	//类型查询
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Println(index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Println(index, value)

		}
	}

}
