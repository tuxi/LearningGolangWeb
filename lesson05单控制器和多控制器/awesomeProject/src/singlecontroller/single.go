package singlecontroller

import (
	"fmt"
	"net/http"
)

/*
单控制器
在Golang的net/http包下有serveMux实现了Front设计模式的Front窗口，serveMux负责接收请求并把请求分发给处理器(handler)
```
type Handler interface {
	ServeHTTP(ResponseWriter *aRequest)
}
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```
* 自定义结构体，实现Handler接口后，这个结构体就属于一个处理器，可以处理全部请求
	* 无论在浏览器中输入的资源地址是什么，都可以访问ServeHTTP
*/

// 实现Handler的结构体MyHandler
type MyHander struct {

}

func (myHander *MyHander)ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World...")
}

// 单控制器事例
func SingleControllerDemo()  {
	fmt.Println("启动单控制器服务")
	myHander := MyHander{}
	server := http.Server{
		Addr: "127.0.0.1:8099",
		Handler: &myHander,
	}
	_ = server.ListenAndServe()
}