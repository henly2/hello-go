package main

import "fmt"

type Vertex4 struct {
	a, b float64
}

var m1 = map[string]Vertex4{
	"zs": {40.68433, -74.39967},
	"ls": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m1)
}
