package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	defer fmt.Println("主携程调用完毕")
	go func() {
		defer fmt.Println("子携程调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子携程 i=", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子携程，"
	}()
	str := <-ch
	fmt.Println("st", str)

}
