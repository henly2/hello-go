package main

import "fmt"

func testaa() {
	fmt.Println("aaa")
}
func testbb(x int) {
	//不想程序崩使用recover
	defer func() {
		//recover()
		//fmt.Println(recover())//打印panic错误
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //调用这个匿名函数

	var a [10]int
	a[x] = 111

	//显示调用panic函数,程序导致中断
	//panic("this is a panic test")
}
func testcc() {
	fmt.Println("ccc")

}
func main() {
	testaa()
	testbb(20)
	testcc()

}
