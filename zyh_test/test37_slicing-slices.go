package main

import "fmt"

func main() {
	p := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("p==", p)
	fmt.Println("p[1:4]==", p[1:4])
	fmt.Println("p[:3]==", p[:3])
	fmt.Println("p[4:]==", p[4:])
}
