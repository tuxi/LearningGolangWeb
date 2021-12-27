package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// 访问：http://127.0.0.1:8090/hello/123 hello后面的路由随便写都可以被访问
	r.HandleFunc("/hello/{key}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Fprintln(writer, "dayinle", vars["key"])
	})
	r.HandleFunc("/abc", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "abc")
	})
	http.ListenAndServe(":8090", r)
}
