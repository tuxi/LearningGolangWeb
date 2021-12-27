package main

import (
	"html/template"
	"net/http"
)

/*
# 一.模版嵌套

* 在实际项目中经常出现页面复用的情况,例如:整个网站的头部信息和底部信息复用
* 可以使用动作{{template "模版名称"}}引用模版
* 引用的模版必须在HTML中定义这个模版
```html
{{define "名称"}}
html
{{end}}
```
*/

func main() {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index.html", "view/head.html", "view/foot.html")
		t.ExecuteTemplate(writer, "layout", nil)
	})
	server.ListenAndServe()
}