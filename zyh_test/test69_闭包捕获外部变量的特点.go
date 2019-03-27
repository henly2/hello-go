package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "zs"
	func() {
		a = 666
		str = "ls"
		fmt.Println(a, str)
	}()
	fmt.Println(a, str)

}
