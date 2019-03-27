package main

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	x, y float64
}

func (v *Vertex7) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	d := &Vertex7{2, 3}
	d.Scale(5)
	fmt.Println(d, d.Abs())

}
