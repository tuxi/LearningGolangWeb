package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func welcome(writer http.ResponseWriter, request *http.Request)  {
	// 解析html文件
	t, _ := template.ParseFiles("view/index.html")
	// 把模版信息写入到输出流中
	// 第二个参数表示向模版传递的数据
	t.Execute(writer, nil)
}

/*
将html输出到浏览器中
在浏览器中输入
```
http://127.0.0.1:8089
```
*/
func demo1()  {
	fmt.Println("加载html模版")
	// 显示index.html信息
	server := http.Server{Addr: ":8089"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

/*
引入静态文件
*/
func demo2()  {
	fmt.Println("加载静态文件")
	server := http.Server{Addr: ":8089"}
	// 访问url以/static/开头，就会把访问信息映射到指定的目录中
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

func main() {
	//demo1()
	demo2()
}
