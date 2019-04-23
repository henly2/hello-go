package util

import (
	"encoding/json"
	"io/ioutil"
)

var dirpath string
func SetSavedDir(dirPath string) {
	dirpath=dirPath
}

func SaveFile(data interface{},fileName string)error{

	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	filename:=dirpath+fileName
	err = ioutil.WriteFile(filename, result, 0666) //写入文件(字符串)
	if err!=nil{
		return err
	}
	return err
}
