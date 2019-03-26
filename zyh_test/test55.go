package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int
	fmt.Println(len(a), len(b))
	a[0] = 1
	for i := 0; i < len(a); i++ {
		a[i] = i + 1

	}
	for i, data := range a {
		fmt.Println(i, data)
	}

}
