package main

import "fmt"

type Students3 struct {
	id   int
	name string
	age  int
}
type Person struct {
	Students3 //匿名字段
	sex       string
}

func main() {
	s1 := Person{Students3{1, "zs", 5}, "nv"}
	fmt.Println(s1)
}
