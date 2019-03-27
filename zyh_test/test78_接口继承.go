package main

import "fmt"

type Humaner0 interface {
	sayhi()
}
type Personer interface {
	Humaner
	sing(lrc string)
}
type Student003 struct {
	name string
	id   int
}

func (tmp *Student003) sayhi() {
	fmt.Println(tmp.name, tmp.id)
}
func (tmp *Student003) sing(lrc string) {
	fmt.Println(lrc)

}
func main() {
	//定义接口类型变量
	var i Personer
	h := &Student003{"mike", 1}
	i = h
	i.sayhi()
	i.sing("kkk")

}
