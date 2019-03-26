package main

import "fmt"

func add01(a, b int) int {
	return a + b

}

type long int

func (tmp long) add02(other long) long {
	return tmp + other

}

func main() {
	result := add01(1, 2)
	fmt.Println(result)
	var a long = 1
	res := a.add02(2)
	fmt.Println(res)

}
