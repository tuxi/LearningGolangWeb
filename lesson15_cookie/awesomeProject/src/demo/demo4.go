package demo

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Expires 设置Cookie 过期时间

func Demo4()  {
	server := http.Server{
		Addr: ":8089",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index3.html")
		t.Execute(writer, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/setCookie", func(writer http.ResponseWriter, request *http.Request) {
		// 验证HttpOnly
		//cookie := http.Cookie{Name: "myKey", Value: "myValue", HttpOnly: false}
		// 验证Path
		//http.Cookie{Name: "myKey", Value: "myValue", Path: "/abc/"}
		// 验证Expirex
		cookie := http.Cookie{Name: "myKey", Value: "myValue", Expires: time.Date(2021, 12, 27, 10, 30, 12, 0, time.Local)}
		http.SetCookie(writer, &cookie)
		t, _ := template.ParseFiles("view/index3.html")
		t.Execute(writer, nil)
	})
	http.HandleFunc("/abc/mypath", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Cookies())
	})
	server.ListenAndServe()
}