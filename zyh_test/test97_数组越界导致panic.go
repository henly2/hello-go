package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}
func testb(x int) {
	var a [10]int
	a[x] = 111

	//显示调用panic函数,程序导致中断
	//panic("this is a panic test")
}
func testc() {
	fmt.Println("ccc")

}
func main() {
	testa()
	testb(20)
	testc()

}
