package main

import "fmt"

type Humaner2 interface {
	sayhi()
}
type Students4 struct {
	name string
	id   int
}

func (tmp *Students4) sayhi() {
	fmt.Println("tmp.name=", tmp.name, "tmp.id=", tmp.id)

}

type Teacher1 struct {
	addr  string
	group int
}

func (tmp *Teacher1) sayhi() {
	fmt.Println("tmp.addr=", tmp.addr, "tmp.group=", tmp.group)

}

func Whosay(i Humaner2) {
	i.sayhi()
}
func main() {
	s := &Students4{"mike", 2}
	t := &Teacher1{"ads", 2}
	Whosay(s)
	Whosay(t)

}
