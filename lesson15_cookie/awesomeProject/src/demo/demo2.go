package demo

import (
	"html/template"
	"net/http"
)

func setCookie1(writer http.ResponseWriter, request *http.Request)  {
	c := http.Cookie{Name: "myKey", Value: "myValue"}
	http.SetCookie(writer, &c)
	t, _ := template.ParseFiles("view/index1.html")
	t.Execute(writer, nil)
}

// Cookie的HttpOnly属性：限制JavaScript中是否可以获取到Cookie
func Demo2()  {
	server := http.Server{
		Addr: ":8089",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建Cookie结构体对象，并设置HttpOnly为true，不允许客户端JavaScript中获取到cookie
		cookie := http.Cookie{Name: "myKey", Value: "myValue", HttpOnly: true}
		http.SetCookie(writer, &cookie)
		t, _ := template.ParseFiles("view/index1.html")
		t.Execute(writer, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/setCookie1", setCookie1)
	server.ListenAndServe()
}

