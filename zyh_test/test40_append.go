package main

import "fmt"

func printslice1(s string, x []int) {
	fmt.Println(s, len(x), cap(x), x)

}
func main() {
	var a []int
	printslice1("a", a)
	a = append(a, 0)
	printslice1("a", a)
	a = append(a, 1)
	printslice1("a", a)
	a = append(a, 2, 3, 4)
	printslice1("a", a)

}
