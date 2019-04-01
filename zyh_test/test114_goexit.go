package main

import (
	"fmt"
	"runtime"
)

func task() {
	defer fmt.Println("cccc")
	runtime.Goexit()
	fmt.Println("dddd")

}
func main() {
	go func() {
		fmt.Println("aaaaaa")
		task()
		fmt.Println("bbbb")
	}()
	for {
		fmt.Println()
		test()
	}

}
