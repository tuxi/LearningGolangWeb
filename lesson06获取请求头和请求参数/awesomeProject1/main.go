package main

import (
	"fmt"
	"net/http"
)

/*
### 获取请求头
*/
func demo1()  {
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}
	http.HandleFunc("/param", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello world...")
		// 获取请求携带的全部请求头
		header := request.Header
		fmt.Fprintln(writer, "Header的全部数据: ", header)

		// 按照名称获取请求头
		encoding := header["Accept-Encoding"]
		fmt.Fprintln(writer, "\n\nAccept-Encoding：", encoding)

		var acc []string = header["Accept"]
		for _, n := range acc {
			fmt.Fprintln(writer, "\n\nAccept:", n)
		}
	})
	server.ListenAndServe()
}

/*
### 获取请求参数
* 请求参数可以一次全部获取，也可以按照名称获取
* 在浏览器地址中输入下面地址，这属于http请求的get方式，请求携带两个参数name和age
```
http://localhost:8090/param?name=yuan&age=18
```
*/
func demo2()  {
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}
	http.HandleFunc("/param", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		// 获取全部的请求参数
		fmt.Fprintln(writer, request.Form) // 输出 map[age:[18] name:[yuan]]
		/*
		按照请求参数名称获取参数值
		根据源码，FormValue(key)=req.Form[key]
		*/
		name := request.FormValue("name")
		age := request.FormValue("age")
		fmt.Fprintln(writer, name, age)
	})
	server.ListenAndServe()
}

func main() {
	//demo1()
	demo2()
}
