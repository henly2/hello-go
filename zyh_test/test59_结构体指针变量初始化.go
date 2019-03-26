package main

import "fmt"

type students1 struct {
	id   int
	name string
	age  int
}

func main() {
	var a *students1 = &students1{1, "zs", 5}
	fmt.Println(a)

}
