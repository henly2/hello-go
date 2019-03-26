package main

import "fmt"

type Person02 struct {
	name string
	sex  byte
	age  int
}

func (p Person02) SetInfovalue() {
	fmt.Println("SetInfovalue")
}
func (p *Person02) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}
func main() {

	p := &Person02{"mike", 'm', 1}
	p.SetInfoPointer()
	p.SetInfovalue()
}
