package main

import "fmt"

type V1 struct {
	X int
	Y int
}

func main() {
	v := V1{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
