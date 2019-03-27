package main

import "fmt"

type IpAdder [4]byte

func main() {
	addrs := map[string]IpAdder{
		"loopback": {127, 0, 0, 1},
		"google":   {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Println(n, a)
	}

}
