package main

import "fmt"

type Person0 struct {
	name string
	sex  byte
	age  int
}

func (p Person0) SetInfoValue() {
	//p.name="mike"
	//p.sex='m'
	//p.age=18
	fmt.Println("SetInfoValue")

}
func (p *Person0) SetInfopointer() {
	fmt.Printf("SetInfopointer")

}
func main() {
	p := &Person0{"mike", 'm', 18}
	p.SetInfopointer()
}
