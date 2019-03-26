package main

import "fmt"

type students struct {
	id   int
	name string
	age  int
}

func main() {
	var a students = students{1, "zs", 5}
	fmt.Println(a)

}
