package main

import "fmt"

type Humaner3 interface {
	sayhi()
}
type Personer2 interface {
	Humaner3
	sing(lrc string)
}
type Students5 struct {
	name string
	id   int
}

func (tmp *Students5) sayhi() {
	fmt.Println(tmp.name, tmp.id)
}
func (tmp *Students5) sing(lrc string) {
	fmt.Println(lrc)
}

func main() {
	var i Personer2
	h := &Students5{"asd", 5}
	i = h
	i.sayhi()
	i.sing("jjj")

}
