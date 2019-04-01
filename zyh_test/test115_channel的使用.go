package main

import (
	"fmt"
	"time"
)

//定义全局变量
var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}

//persom1执行完，才能到person2执行
func person1() {
	Printer("hello")
	ch <- 666 //给管道写数据，发送
}
func person2() {
	<-ch //从管道取数据，接收，如果通道没有数据就会阻塞
	Printer("world")

}
func main() {
	go person1()
	go person2()
	for {
		fmt.Println()
	}

}
