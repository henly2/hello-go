package main

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v, lim)
	}
	return lim
}
func main() {
	fmt.Println(pow1(5, 2, 3), pow1(2, 3, 1))

}
