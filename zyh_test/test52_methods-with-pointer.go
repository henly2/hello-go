package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	x, y float64
}

func (v *Vertex3) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func (v *Vertex3) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := &Vertex3{3, 4}
	v.Scale(5)
	fmt.Println(v, v.abs())

}
