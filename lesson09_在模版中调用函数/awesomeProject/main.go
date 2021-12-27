package main

import (
	"html/template"
	"net/http"
	"time"
)

/*
### 在模版中调用函数
* 在模版中调用函数时,如果是无参函数直接调用函数名即可,没有函数的括号
* 例如在go源码中`时间变量.Year()`在模版中`{{时间.Year}}`
* 在模版中调用有参函数时参数和函数名称之间有空格,参数和参数之间也是空格
*/

/*
调用系统函数/方法
*/
func demo1() {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index.html")
		time := time.Date(2021,12,25,16,38,10,0, time.Local)
		t.Execute(writer, time)
	})
	server.ListenAndServe()
}

/*
调用自定义函数/方法
* 如果希望调用z自定义函数，需要借助html/template包下的FuncMap进行映射
* FUncMap本质上就是map的别名`type FuncMap map[string]interface{}`
* 函数被添加映射后，只能通过在FuncMap中key调用函数
*/

func MyFormat(t time.Time) string {
	return t.Format("2008-01-19 15:03:36")
}

func demo2()  {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 把自定义函数绑定到FFuncMap上
		funcMap := template.FuncMap{"mf": MyFormat}
		// 此处注意，一定要先绑定函数
		t := template.New("index1.html").Funcs(funcMap)
		// 绑定函数后，解析模版
		t, _ = t.ParseFiles("view/index1.html")
		time1 := time.Date(2018,1,2,3,4,5,0,time.Local)
		t.Execute(writer, time1)

	})
	server.ListenAndServe()
}

func main() {
	//demo1()
	demo2()
}