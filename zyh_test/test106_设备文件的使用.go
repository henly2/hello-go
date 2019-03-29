package main

import (
	"fmt"
)

func main() {
	//os.Stdout.Close()//关闭文件
	//fmt.Println("are you ok?")
	//os.Stdout.WriteString("are you ok")
	var a int
	fmt.Println("请输入a")
	fmt.Scan(&a)
	fmt.Println("a=", a)
}
