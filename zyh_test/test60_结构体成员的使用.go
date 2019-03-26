package main

import "fmt"

type students2 struct {
	id   int
	name string
	age  int
}

func main() {
	var s students2
	s.id = 1
	s.age = 5
	s.name = "zs"
	fmt.Println(s)

}
