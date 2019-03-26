package main

import "fmt"

type vertex3 struct {
	a, b float64
}

var m2 = map[string]vertex3{
	"zs": vertex3{1.2, 2.3},
	"ls": vertex3{5.4, 4.8},
}

func main() {
	fmt.Println(m2)

}
