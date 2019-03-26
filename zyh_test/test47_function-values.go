package main

import (
	"fmt"
	"math"
)

func main() {
	a := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)

	}
	fmt.Println(a(1, 2))

}
