package main

import "fmt"

func printslice(s string, x []int) {
	fmt.Println(s, len(x), cap(x), x)

}
func main() {
	a := make([]int, 5)
	printslice("a", a)
	b := make([]int, 0, 5)
	printslice("b", b)
	c := b[:2]
	printslice("c", c)
	d := c[2:5]
	printslice("d", d)

}
