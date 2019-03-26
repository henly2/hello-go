package main

import "fmt"

type Vertex2 struct {
	a, b float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.1, -20.5,
	}
	fmt.Println(m["Bell Labs"])

}
