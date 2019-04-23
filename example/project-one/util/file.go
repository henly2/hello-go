package util

import (
	"encoding/json"
	"io/ioutil"
)

var (
	savedDir = "./"
)

func SetSavedDir(dir string) {
	savedDir = dir

	// TODO:可以进一步检查目录是否存在，如果不存在，返回错误
}

func SaveData(data interface{}, fileName string) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	filePath := savedDir + fileName
	err = ioutil.WriteFile(filePath, dataBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadData(fileName string, data interface{}) error {
	filePath := savedDir + fileName
	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, data)
	if err != nil {
		return err
	}

	return nil
}