package util

//代码抽取

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//读取并得到行情路径
func ReadFile(path1 string) ([]byte, error) {
	//return ioutil.ReadFile(path1)

	//打开文件
	f, err := os.Open(path1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	buf1 := make([]byte, 1024*2)
	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf1)
	if err1 != nil {
		fmt.Println(err1)
		return nil, err
	}
	fmt.Println("打印buf1=", string(buf1[:n]))
	return buf1[:n], nil

}

//得到拼接后的路径
func Req(data []byte) (fi_path1 string) {
	type Cfg struct {
		Url    string `json:"url"`
		Symbol string `json:"symbol"`
	}
	m := Cfg{}
	err5 := json.Unmarshal([]byte(data), &m) //第二个参数需要地址传递
	if err5 != nil {
		fmt.Println("err5=", err5)
		return
	}
	fmt.Println("m.,Url=", m.Url, "m.Symbol=", m.Symbol)

	file_path := m.Url + "/market/detail/merged?symbol=" + m.Symbol
	fmt.Println(file_path)
	return file_path
}

//通过行情路径读取行情数据并写入到指定文件中
func Http_data(file_path string) {
	resp, err := http.Get(file_path)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	//fmt.Println(string(body))

	//生成json文本
	//m:=make(map[string]interface{},4)
	result, err1 := json.Marshal(body)
	if err != nil {
		fmt.Println(err1)
		return
	}
	//读取写入的数据
	fmt.Println(string(result))

	//写入到文件
	fi, error1 := os.Create("./data.json")
	if error1 != nil {
		fmt.Println(error1)
	}
	data := string(result)
	for i := 0; i < 5000; i++ {
		//写入byte的slice数据
		//fi.Write([]byte(data))
		//写入字符串
		fi.WriteString(data)
	}
	fi.Close()

}
