package main

import "fmt"

type Vertex1 struct {
	x, y int
}

func main() {
	v1 := Vertex1{1, 2}
	v2 := Vertex1{x: 1}
	v3 := Vertex1{}
	p := &Vertex1{1, 2}
	fmt.Println(v1, v2, v3, *p)
}
