package main

import (
	"goweb/framework"
	"net/http"
)

//
func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: framework.NewCore(),
		Addr:    ":8080", // 请求监听地址
	}
	server.ListenAndServe()

	fs := http.FileServer(http.Dir(""))
	http.Handle("", http.StripPrefix("", fs))
}
