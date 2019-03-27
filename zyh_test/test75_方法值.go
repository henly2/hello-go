package main

import "fmt"

type Person002 struct {
	name string
	sex  byte
	age  int
}

func (p *Person002) SetInfopoint() {
	fmt.Println(p)

}
func (p Person002) SetInfovalue() {
	fmt.Println(p)

}
func main() {
	p := Person002{"mike", 'm', 12}
	p.SetInfopoint()
	pFunc := p.SetInfopoint //方法值
	pFunc()

}
