package main

import (
	"fmt"
	"net/http"
)

func welome(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(writer, "收到消息<br>Hello world...</br>")
}

func main() {
	http.HandleFunc("/", welome)
	// 监听本机端口8081
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println("监听服务失败", err)
	}
}
