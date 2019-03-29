package main

import (
	"./util"
	"fmt"
)

/*
功能：从火币网获取数字火币的行情数据，并将数据保存到文件

1. 配置文件放在cfg/cfg1.json文件，从里面获取url和symbol

2. 获取行情
3. 将行情数据保存到data/data.json文件

火币行情: https://api.huobi.pro/market/detail/merged?symbol=btcusdt
文档地址: https://huobiapi.github.io/docs/v1/cn
*/

func main() {
	//获取行情数据地址
	path := "src/github.com/henly2/hello-go/test/cfg1/cfg1.json"
	//返回读取到的内容data接收

	data, err := util.ReadFile(path)

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//fmt.Println("返回读取到的值：",string(data))

	//得到拼接好的行情路径
	file_path := util.Req(data)
	//通过拼接的地址进行获取数据
	util.Http_data(file_path)

}
