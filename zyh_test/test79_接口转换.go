package main

import "fmt"

type Humaner1 interface {
	sayhi()
}
type Personer1 interface {
	Humaner
	sing(lrc string)
}
type Student004 struct {
	name string
	id   int
}

func (tmp *Student004) sayhi() {
	fmt.Println(tmp.name, tmp.id)
}
func (tmp *Student004) sing(lrc string) {
	fmt.Println(lrc)

}
func main() {
	//定义接口类型变量
	//var i Personer
	//h:=&Student004{"mike", 1}
	//i=h
	//i.sayhi()
	//i.sing("kkk")
	var P Personer
	var i Humaner1
	i = P
	i.sayhi()

}
