package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("nil")
	}

}
