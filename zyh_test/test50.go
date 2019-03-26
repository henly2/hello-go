package main

import (
	"fmt"
)

type myInt int

func Add(a, b int) { //函数
	fmt.Println(a + b)
}

func (a myInt) Add(b myInt) { //方法
	fmt.Println(a + b)
}

func main() {
	a, b := 3, 4
	var aa, bb myInt = 3, 4
	Add(a, b)
	aa.Add(bb)
}
