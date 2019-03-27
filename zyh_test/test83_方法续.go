package main

import "fmt"

type M1 float64

func (f M1) abs() float64 {
	if f < 0 {
		return float64(f)

	}
	return float64(-f)

}
func main() {
	b := M1(-2)
	fmt.Println(b.abs())

}
