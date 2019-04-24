package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("开始http服务,端口8080...")
	log.Println("http://127.0.0.1:8080/")
	log.Println("http://127.0.0.1:8080/api")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I am index")) //设置访问的路由
	})

	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I am api"))
	})

	http.ListenAndServe(":8080", nil) //设置监听的端口
}
