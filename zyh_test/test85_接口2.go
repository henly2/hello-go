package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex8 struct {
	x, y float64
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.y * v.x)
}
func main() {
	var a Abser
	f := MyFloat(math.Sqrt(2))
	v := &Vertex8{2, 3}
	a = f
	fmt.Println(a.Abs())

	a = v
	fmt.Println(a.Abs())

}
