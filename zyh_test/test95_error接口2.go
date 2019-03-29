package main

import (
	"errors"
	"fmt"
)

func Mydiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return result, err

}
func main() {
	//Mydiv(4,5)
	//fmt.Println(Mydiv(4,5))
	result, err := Mydiv(10, 5)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("result", result)
	}
}
