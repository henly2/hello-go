package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	x, y float64
}

func (v *Vertex6) ABS() float64 {
	return math.Sqrt(v.x * v.y)

}
func main() {
	a := &Vertex6{2, 3}
	fmt.Println(a.ABS())

}
