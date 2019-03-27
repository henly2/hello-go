package main

import "fmt"

type Person003 struct {
	name string
	sex  byte
	age  int
}

func (p *Person003) SetInfopoint() {
	fmt.Println(p)

}
func (p Person003) SetInfovalue() {
	fmt.Println(p)

}
func main() {
	p := Person003{"mike", 'm', 12}
	p.SetInfopoint()
	pFunc := p.SetInfopoint //方法值
	pFunc()
	f := (*Person003).SetInfopoint
	f(&p)

}
