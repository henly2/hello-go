package main

import "fmt"

func test1() {
	fmt.Println("aaa")
}
func test2() {
	fmt.Println("bbb")
	//显示调用panic函数,程序导致中断
	panic("this is a panic test")
}
func test3() {
	fmt.Println("ccc")

}
func main() {
	test1()
	test2()
	test3()

}
