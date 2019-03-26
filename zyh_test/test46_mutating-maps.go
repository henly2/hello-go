package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key1"] = 42
	fmt.Println("the value", m["key1"])
	m["key2"] = 43
	fmt.Println("the value", m["key2"])
	//delete(m,"key1")
	//fmt.Println("the value",m["key1"])
	fmt.Println(m)
	v, ok := m["key1"]
	fmt.Println("the value", v, "present?", ok)

}
