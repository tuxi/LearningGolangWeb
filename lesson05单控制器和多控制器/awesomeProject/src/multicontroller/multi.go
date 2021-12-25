package multicontroller

import (
	"fmt"
	"net/http"
)

/*
多控制器
* 在实际开发中大部分情况不应该只有一个控制器，不同的请求应该交给不同的处理但愿，在Golang中支持两种多处理方式
	* 多个处理器(Handler)
	* 多个处理函数(HandleFunc)
* 使用多处理器
	* 使用http.Handle把不同的URL绑定到不同的处理器
	* 在浏览器输入http://localhost:8089/myhandler或http://localhost:8089/myother可以访问两个处理器方法，但是访问其他URL会出现404（资源未找到）页面
*/

// 定义实现Handle接口的结构体
type MyHandler struct {}
type MyOtherHandler struct {}

func (handler *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello my handler")
}

func (hander *MyOtherHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello my other handler")
}

// 多handler，需要定义实现Handle接口的结构体
func MultiControllerDemo1()  {
	fmt.Println("启动多控制器服务，多处理器的方式")
	myHandler := MyHandler{}
	myOtherHandler := MyOtherHandler{}
	http.Handle("/myhandler", &myHandler)
	http.Handle("/myother", &myOtherHandler)
	server := http.Server{
		Addr: "localhost:8089",
	}
	server.ListenAndServe()
}

// 多函数
// 处理方式要比多处理器简便，直接把资源路径与函数绑定

func first(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "第一个")
}

//func second(res http.ResponseWriter, req *http.Request)  {
//	fmt.Fprintln(res, "第二个")
//}

func MultiControllerDemo2() {
	// 定义函数接送回调
	http.HandleFunc("/first", first)
	// 使用匿名函数回调
	http.HandleFunc("/second", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "第二个")
	})
	server := http.Server{
		Addr: "127.0.0.1:8089",
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("启动服务失败", err)
	}
}