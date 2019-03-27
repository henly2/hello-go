package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Student002 struct {
	name string
	id   int
}

func (tmp *Student002) sayhi() {
	fmt.Println(tmp.name, tmp.id)

}

type Teacher struct {
	addr  string
	group int
}

func (tmp *Teacher) sayhi() {
	fmt.Println(tmp.addr, tmp.group)

}

type Mystr string

func (tmp *Mystr) sayhi() {
	fmt.Println(tmp)

}
func WhoSayhi(i Humaner) {
	i.sayhi()

}
func main() {
	s := &Student002{"mike", 3}
	t := &Teacher{"bj", 1}
	var str Mystr
	WhoSayhi(s)
	WhoSayhi(t)
	WhoSayhi(&str)

}

//func main() {
//	//定义接口类型的变量
//	var i Humaner
//	s:=&Student002{"mike", 1}
//	i=s
//	i.sayhi()
//	t:=&Teacher{"mike", 1}
//	i=t
//	i.sayhi()
//	var str Mystr ="hello"
//	i=&str
//	i.sayhi()
//
//
//
//}
