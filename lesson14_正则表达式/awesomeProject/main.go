package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func main() {
	//demo1()
	demo2()
}

// 正则匹配
func demo1()  {
	// 创建结构体变量，匹配数字和字母组合
	r := regexp.MustCompile(`\d[a-zA-Z]`)
	// 判断是否匹配
	fmt.Println(r.MatchString("6A2")) // 输出 true
	/*
	字符串中满足要求的片段，返回[]string
	第二个参数是[]string的长度，-1表示不限长度
	*/
	fmt.Println(r.FindAllString("56A6B7C", -1)) // 输出 [6A 6B 7C]
	/*
	把正则表达式匹配的结构当作分隔符，拆分字符串
	n > 0: 返回最多n个字符串，最后一个字符串是剩余未分割的部分
	n == 0: 返回nil（zero substrings）
	n < 0: 返回所有字符串
	*/
	fmt.Println(r.Split("12345qwert", -1)) // 输出 [1234 wert]
	// 把满足正则要求的内容替换为指定字符串
	fmt.Println(r.ReplaceAllString("12345qwert", "替换的字符")) // 输出 1234替换的字符wert
}

func demo2()  {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index.html")
		t.Execute(writer, nil)
	})
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		username := request.FormValue("username")
		r, _ := regexp.MatchString(`^[0-9a-zA-Z]]{6,12}$`, username)
		if r {
			fmt.Fprintln(writer, "注册成功")
		} else {
			fmt.Fprintln(writer, "用户格式不正确")
		}
	})
	server.ListenAndServe()
}