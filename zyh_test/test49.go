package main

import (
	"fmt"
	"math"
)

type a struct {
	x, y float64
}

func (v a) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)

}
func main() {
	v := a{2, 3}
	fmt.Println(v.Abs())

}
