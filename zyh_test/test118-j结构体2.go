package main

import "fmt"

type Students struct {
	id   int
	name string
	age  int
}

func main() {
	var a Students = Students{3, "zs", 9}
	fmt.Println(a)

}
