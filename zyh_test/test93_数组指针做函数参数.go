package main

import "fmt"

func modify1(a *[5]int) {
	a[0] = 9
	fmt.Println(a)

}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify1(&a)
	fmt.Println(a)

}
