package main

import (
	"html/template"
	"net/http"
)

func myFunc(writer http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	m := map[string]string{"key1": "value1", "key2": "value2"}
	t.Execute(writer, m)
}

func main() {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", myFunc)
	server.ListenAndServe()
}
