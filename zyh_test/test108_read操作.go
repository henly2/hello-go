package main

import (
	"fmt"
	"os"
)

func ReadFile1(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2)
	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("buf=", string(buf[:n]))
}

func main() {
	path := "./demo.txt"
	ReadFile1(path)

}
