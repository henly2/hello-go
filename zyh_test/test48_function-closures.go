package main

import "fmt"

func add1() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum

	}

}
func main() {
	a, b := add1(), add1()
	for i := 0; i < 10; i++ {
		fmt.Println(a(2), b(3))

	}

}
