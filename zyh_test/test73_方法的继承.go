package main

import "fmt"

type Person00 struct {
	name string
	sex  byte
	age  int
}

func (tmp *Person00) PointInfo() {
	fmt.Println(tmp.name, tmp.sex, tmp.age)

}

type Student00 struct {
	Person00
	id   int
	addr string
}

func main() {
	s := Student00{Person00{"mike", 'm', 18}, 5, "aaa"}
	s.PointInfo()
}
