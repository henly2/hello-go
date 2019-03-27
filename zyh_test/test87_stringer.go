package main

import "fmt"

type Person3 struct {
	name string
	age  int
}

func (p Person3) String() string {
	return fmt.Sprintf("%v(%v years)", p.name, p.age)

}
func main() {
	a := Person3{"mike", 2}
	z := Person3{"amy", 5}
	fmt.Println(a, z)

}
