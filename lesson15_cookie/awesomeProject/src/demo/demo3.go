package demo

import (
	"fmt"
	"html/template"
	"net/http"
)

// Path 属性设置Cookie的访问范围
func Demo3()  {
	server := http.Server{
		Addr: ":8089",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index2.html")
		t.Execute(writer, nil)
	})
	http.HandleFunc("setCookie", func(writer http.ResponseWriter, request *http.Request) {
		// 验证HttpOnly
		//cookie := http.Cookie{Name: "myKey", Value: "myValue", HttpOnly: false}
		// 验证Path
		cookie := http.Cookie{Name: "myKey", Value: "myValue", Path: "/abc/"}
		http.SetCookie(writer, &cookie)
		t, _ := template.ParseFiles("view/index2.html")
		t.Execute(writer, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// 客户端想要获取Cookie，路径必须在/abc/下面的
	http.HandleFunc("/abc/mypath", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Cookies())
	})
	server.ListenAndServe()
}