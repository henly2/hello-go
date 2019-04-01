package main

import "fmt"

type Students1 struct {
	id   int
	name string
	age  int
}

func main() {
	var s Students1
	s.id = 4
	s.name = "b"
	s.age = 5
	fmt.Println("s=", s)

}
