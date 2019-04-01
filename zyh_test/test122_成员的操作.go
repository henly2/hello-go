package main

import "fmt"

type Person4 struct {
	name string
	sex  string
	age  int
}
type Students2 struct {
	Person4
	id   int
	addr string
}

func main() {
	s1 := Students2{Person4{"zs", "f", 5}, 1, "ccc"}
	s1.name = "e"
	s1.age = 6
	s1.sex = "m"
	fmt.Println(s1.name, s1.sex, s1.age)

}
