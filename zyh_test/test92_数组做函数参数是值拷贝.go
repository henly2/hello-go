package main

import "fmt"

func modify(a [5]int) {
	a[0] = 9
	fmt.Println("modify a=", a)

}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println("main a=", a)

}
