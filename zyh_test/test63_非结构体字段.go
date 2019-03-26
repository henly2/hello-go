package main

import "fmt"

type Myster string //给类型改名
type Person2 struct {
	name string
	sex  string
	age  int
}
type Student1 struct {
	Person2
	int
	Myster
}

func main() {
	s1 := Student1{Person2{"zs", "m", 18}, 1, "ss"}
	fmt.Println(s1)

}
