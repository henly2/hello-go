package main

import "fmt"

type Person01 struct {
	name string
	sex  byte
	age  int
}

func (tmp Person01) PeintInfo() {
	fmt.Println(tmp)

}
func (p *Person01) SetInfo(a string, b byte, c int) {
	p.name = a
	p.sex = b
	p.age = c

}
func main() {
	p := Person01{"ls", 'm', 6}
	p.PeintInfo()

	var p2 Person01
	(&p2).SetInfo("zs", 'f', 5)
	p.PeintInfo()

}
