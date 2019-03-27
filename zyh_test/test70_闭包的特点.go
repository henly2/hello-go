package main

import "fmt"

func test() func() int {
	var i int
	return func() int {
		i++
		return i * i
	}

}
func main() {
	f := test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
