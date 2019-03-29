package main

import (
	"fmt"
	"strings"
)

func main() {
	//看“hellogo”是否包含“he"
	fmt.Println(strings.Contains("hellogo", "he"))
	s := []string{"abc", "hello", "go"}
	buf := strings.Join(s, "$") //用"$连接字符串
	fmt.Println(buf)

}
