package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x.")
	case "linux":
		fmt.Println("linux")
	default:

		fmt.Println("%s", os)
	}

}
