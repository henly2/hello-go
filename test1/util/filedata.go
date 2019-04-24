package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var dirpath string

func SetSavedDir(dirPath string) {
	dirpath = dirPath
}

func SaveFile(data interface{}, fileName string) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	filename := dirpath + fileName
	err = ioutil.WriteFile(filename, result, 0666) //写入文件(字符串)
	if err != nil {
		return err
	}
	return err
}
func ReadFile(data interface{}, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		filedata, err := ioutil.ReadAll(f)
		if err != nil {
			return
		}
		if data == string(filedata) {
			fmt.Println("输入数据成功")
		} else {
			fmt.Println("输入数据失败")
		}
		fmt.Println("fileddata=", string(filedata), "data=", data)

	}

}
