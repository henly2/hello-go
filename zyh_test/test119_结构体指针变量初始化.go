package main

import "fmt"

type Students0 struct {
	id   int
	name string
	age  int
}

func main() {
	var a *Students0 = &Students0{5, "a", 5}
	fmt.Println(a)
}
