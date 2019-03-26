package main

import (
	"fmt"
)

type M float64

func (f M) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	var ff M = 3
	f := M(2)
	fmt.Println(ff.abs())
	fmt.Println(f.abs())
}
