package demo

import (
	"fmt"
	"html/template"
	"net/http"
)

func welcome(writer http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(writer, nil)
}

func setCookie(writer http.ResponseWriter, request *http.Request)  {
	c := http.Cookie{Name: "myKey", Value: "myValue"}
	http.SetCookie(writer, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(writer, nil)
}

func getCookie(writer http.ResponseWriter, request *http.Request)  {
	// 根据key取出Cookie
	//c := request.Cookie("myKey")
	// 取出全部Cookie
	cookies := request.Cookies()
	for _, n := range cookies{
		fmt.Println(n.Name, n.Value)
	}
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(writer, cookies)
}


// 生成Cookie和获取Cookie
func Demo1()  {
	server := http.Server{
		Addr: ":8089",
	}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("sttic"))))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}