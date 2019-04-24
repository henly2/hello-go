package main

import (
	"encoding/json"
	"fmt"
	"github.com/henly2/hello-go/test1/util"
)

func main() {
	dirPath := "./"
	data := "abcdefgkkk"
	fileName := "datafile.txt"
	util.SetSavedDir(dirPath)
	util.SaveFile(data, fileName)
	result, _ := json.Marshal(data)
	datafile := string(result)
	util.ReadFile(datafile, fileName)
	fmt.Println("main_data=", datafile)
}
