//代码抽取
package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//读取并得到行情路径
func ReadFile(path string) ([]byte, error) {
	//return ioutil.ReadFile(path1)
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	buf1 := make([]byte, 1024*2)
	n, err := f.Read(buf1) //n代表从文件读取内容的长度
	if err != nil {
		fmt.Println(err)
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
	err := json.Unmarshal([]byte(data), &m) //第二个参数需要地址传递
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("m.,Url=", m.Url, "m.Symbol=", m.Symbol)

	file_path := m.Url + "/market/detail/merged?symbol=" + m.Symbol
	fmt.Println(file_path)
	return file_path
}

//读取行情数据并写入到指定文件中
func Http_data(file_path string) (result []byte) {
	resp, err := http.Get(file_path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	result, err1 := json.Marshal(body)
	if err != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(result)) //读取写入的数据
	return result

}

//写入到./data.json文件中
func Write(result []byte) {
	fmt.Println(result)
	//写入到文件
	fi, error1 := os.Create("./data.json")
	if error1 != nil {
		fmt.Println(error1)
	}
	for i := 0; i < 5000; i++ {
		//写入byte的slice数据
		fi.Write([]byte(result))
		//写入字符串
		//fi.WriteString(result);
	}
	fi.Close()
}
