package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	x, y float64
}

func (v *Vertex5) abs() float64 {
	return math.Sqrt(v.x + v.y)

}

type Myfloat float64

func (f Myfloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type abser interface {
	abs() float64
}

func main() {
	var aa abser
	f := Myfloat(-math.Sqrt(2))
	v := Vertex5{1, 2}
	aa = f
	aa = &v
	fmt.Println(aa.abs())

}
