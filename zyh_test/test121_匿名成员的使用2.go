package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "aaa"
	f1 := func() {
		fmt.Println(a)
		fmt.Println(str)
	}
	f1()

	//带参数的匿名函数
	f2 := func(i, j int) {
		fmt.Println(i, j)
	}
	f2(10, 20)

	func(i, j int) {
		fmt.Println(i, j)
	}(10, 20)

	//有参数有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(60, 20)
	fmt.Println(x, y)
}
