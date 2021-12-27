package main

import (
	"html/template"
	"net/http"
)

/*
事例1
向模版中传递字符串
* 向HTML传递字符串数据.在HTML中使用{{.}}获取传递数据即可.所有基本类型都是使用此方式进行传递
```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
</head>
<body>
<pre>
尊敬的{{.}}先生/女士
    您已经被我公司录取,收到此消息后请您仔细阅读附件中"注意事项"
    再次祝您:{{.}}好运
</pre>
</body>
</html>
```
*/
func demo1()  {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/first", welcome)
	server.ListenAndServe()
}
func welcome(writer http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("view/index1.html")
	// 向模版中传递字符串
	t.Execute(writer, "杨校园")
}

/*
事例2
向模版中传递结构体类型数据
* 结构体中的属性首字母必须大写才能被模版访问
* 在模版中直接使用`{{.属性名}}`获取结构体的属性
*/

type User struct {
	Name string
	Age int
}

func demo2()  {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/second", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index.html")
		t.Execute(writer, User{"远远", 22})
	})
	server.ListenAndServe()
}

/*
事例3
向模版传递map类型数据
* 只能使用`{{.key}}`获取map中的数据
* 模版中支持连缀写法(不仅仅是map)
*/
func demo3()  {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index2.html")
		// 定义一个字典，将其传递到模版中
		m := make(map[string]interface{})
		m["user"] = User{"杨远", 18}
		m["eth"] = 100001
		t.Execute(writer, m)
	})
	server.ListenAndServe()
}

func main() {
	//demo1()
	//demo2()
	demo3()
}

