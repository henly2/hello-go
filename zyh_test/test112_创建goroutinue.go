package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newtask")
		time.Sleep(time.Second)
	}

}
func main() {
	go newTask() //创建一个携程

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}
}
