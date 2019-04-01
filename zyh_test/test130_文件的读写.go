package main

import (
	"fmt"
	"os"
)

func WriteFile1(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		f.WriteString(buf)
	}

}
func main() {
	path := "./demo1.txt"
	WriteFile1(path)

}
