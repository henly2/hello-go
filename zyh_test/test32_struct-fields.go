package main

import "fmt"

type V struct {
	x int
	y int
}

func main() {
	v := V{1, 2}
	v.x = 5
	fmt.Println(v.x)

}
