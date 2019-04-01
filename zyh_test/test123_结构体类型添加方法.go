package main

import "fmt"

type Person5 struct {
	name string
	sex  byte
	age  int
}

func (tmp Person5) PaintInfo() {
	fmt.Println("tmp=", tmp)
}

func (p *Person5) SetInfo(a string, b byte, c int) {

	p.name = a
	p.sex = b
	p.age = c
}
func main() {
	p := Person5{"s", 'm', 5}
	p.PaintInfo()
	var p2 Person5
	(&p2).SetInfo("d", 'm', 6)
	p2.PaintInfo()

}
