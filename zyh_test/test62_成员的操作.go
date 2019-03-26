package main

import "fmt"

type Person1 struct {
	name string
	sex  string
	age  int
}
type Student struct {
	Person1
	id   int
	addr string
}

func main() {
	s1 := Student{Person1{"zs", "m", 18}, 1, "ss"}
	s1.name = "yoyo"
	s1.sex = "f"
	s1.age = 2
	fmt.Println(s1.name, s1.sex, s1.age, s1.addr)
}
