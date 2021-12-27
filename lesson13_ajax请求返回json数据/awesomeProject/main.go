package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

/*
# JSON简介

* 轻量级数据传输**格式**

* 总体上分为两种:

  * 一种是JSONObject(json对象)

  ```
  {"key":value,"key":value}
  ```

  * 一种是JSONArrayP(json数组),包含多个JSONObject

  ```
  [{"key":"value"},{"key":"value"}]
  ```

* key是string类型,value可以是string类型(值被双引号包含),也可以是数值或布尔类型等,也可以是JSONObject类型或JSONArray类型

* 可以使用Go语言标准库中 encoding/json 包下的Marshal()或Unmarshal()把结构体对象转换成[]byte或把[]byte中信息写入到结构体对象中

  * 在转换过程中结构体属性tag中定义了json中的key,属性的值就是json中的value
  * 如果属性没有配置tag,属性就是json中的key

* 属性的tag可以进行下面配置

```go
// 字段被本包忽略
Field int `json:"-"`
// 字段在json里的键为"myName"
Field int `json:"myName"`
// 字段在json里的键为"myName"且如果字段为空值将在对象中省略掉
Field int `json:"myName,omitempty"`
// 字段在json里的键为"Field"（默认值），但如果字段为空值会跳过；注意前导的逗号
Field int `json:",omitempty"`
```

*/

type User struct {
	Name string
	Age int
}

func main() {
	//demo1()
	demo2()
}

// 对象转json
func demo1()  {
	// 将结构体转json
	u := User{"yuan", 18}
	// 将user对象转换为字节切片
	b, _ := json.Marshal(u)
	fmt.Println(b)
	// 把字节切片转换为user对象
	u1 := new(User)
	json.Unmarshal(b, u1)
	fmt.Println(u1)

}

/*
Ajax访问返回json数据
* 使用jquery封装的$.post()进行ajax请求
* HTML页面发送ajax请求，请求数据
*/
func demo2()  {
	server := http.Server{
		Addr: ":8090",
	}
	// 首页
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("view/index.html")
		t.Execute(writer, nil)
	})
	// 获取用户信息
	http.HandleFunc("/getUser", func(writer http.ResponseWriter, request *http.Request) {
		users := make([]User, 0)
		users = append(users, User{"张三", 12})
		users = append(users, User{"王四",18})
		users = append(users, User{"李五", 20})
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		b, _ := json.Marshal(users)
		fmt.Fprintln(writer, string(b))
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.ListenAndServe()
}