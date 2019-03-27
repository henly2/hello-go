package main

import "fmt"

type Person001 struct {
	name string
	sex  byte
	age  int
}

func (tmp *Person001) PointInfo() {
	fmt.Println(tmp.name, tmp.sex, tmp.age)

}

type Student001 struct {
	Person001
	id   int
	addr string
}

func (tmp *Student001) PointInfo() {
	fmt.Println(tmp)
}
func main() {
	s := Student001{Person001{"mike", 'm', 18}, 5, "aaa"}
	s.PointInfo()
}
