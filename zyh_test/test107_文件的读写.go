package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//使用完毕关闭文件
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i) //把i写入到buf中
		f.WriteString(buf)
	}

}
func main() {
	path := "./demo.txt"
	WriteFile(path)

}
