package main

import "fmt"

func main() {
	var a [10]int
	a[0] = 1
	a[1] = 2
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
		fmt.Printf("i=%d,a[i]=%d\n", i, a[i])
	}
	for i, data := range a {
		fmt.Println(i, data)
	}

}
