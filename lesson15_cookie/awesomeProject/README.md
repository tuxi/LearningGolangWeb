### Cookie 简介

* Cookie就是客户端存储技术.以键值对的形式存在
* 在B/S架构中,服务器端产生Cookie响应给客户端,浏览器接收后把Cookie存在在特定的文件夹中,以后每次请求浏览器会把Cookie内容放入到请求中

### Go语言对Cookie的支持

* 在net/http包下提供了Cookie结构体

    * Name设置Cookie的名称
    * Value 表示Cookie的值
    * Path 有效范围
    * Domain 可访问Cookie 的域
    * Expires 过期时间
    * MaxAge 最大存活时间,单位秒
    * HttpOnly 是否可以通过脚本访问

  ```go
  type Cookie struct {
  	Name  string
  	Value string

  	Path       string    // optional
  	Domain     string    // optional
  	Expires    time.Time // optional
  	RawExpires string    // for reading cookies only

  	// MaxAge=0 means no 'Max-Age' attribute specified.
  	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
  	// MaxAge>0 means Max-Age attribute present and given in seconds
  	MaxAge   int
  	Secure   bool
  	HttpOnly bool
  	Raw      string
  	Unparsed []string // Raw text of unparsed attribute-value pairs
  }
  ```

### 代码演示

* 默认显示index.html页面,显示该页面时没有Cookie,点击超链接请求服务器后,服务端把Cookie响应给客户端,通过开发者工具(F12)观察整个过程.

  ```go
  <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
          "http://www.w3.org/TR/html4/loose.dtd">
  <html>
  <head>
      <title></title>
  </head>
  <body>
  <a href="setCookie">产生Cookie</a>
  <a href="getCookie">获取Cookie</a>
  <br/>
  {{.}}
  </body>
  </html>
  ```

* 服务器提供创建Cookie和获取Cookie的代码
```go
package main

import (
	"net/http"
	"html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "mykey", Value: "myvalue"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)

}
func getCookie(w http.ResponseWriter, r *http.Request) {
	//根据key取出Cookie
	//c1,_:=r.Cookie("mykey")
	//取出全部Cookie内容
	cs := r.Cookies()
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, cs)
}
func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}

```

### HttpOnly
* HttpOnly 控制Cookie的内容是否可以被JavaScript 访问到，通过设置HttpOnly为true时防止XSS攻击防御手段之一
* 默认HttpOnly为false，表示客户端可以通过js获取
* 在项目中导入jquery.cookie.js库，使用jquery获取客户端Cookie内容
* HTML代码如下

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="/static/js/jquery-1.7.2.js"></script>
    <script src="/static/js/jquery.cookie.js"></script>
    <script type="text/javascript">
        $(function () {
            $("button").click(function () {
                var value = $.cookie("mykey")
                alert(value)
            })
        })
    </script>
</head>
<body>
<a href="setCookie">产生Cookie</a>
<button>获取cookie</button>
</body>
</html>
```

* 服务端代码如下

```go
package main

import (
	"net/http"
	"html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "mykey", Value: "myvalue", HttpOnly: false}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)

}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	server.ListenAndServe()
}

```

### Path
* Path属性设置Cookie的访问范围
* 默认为"/"表示当前项目下所有都可以访问
* Path设置路口及子路径内容都可以访问
* 首先先访问index.html，点击超链接产生cookie，在浏览器地址输入localhost:8090/abc/mypath后发现可以访问cookie
* html代码没有变化，只需要修改服务器代码如下

```go
package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	//验证httponly
	//c := http.Cookie{Name: "mykey", Value: "myvalue", HttpOnly: false}
	//验证path
	c := http.Cookie{Name: "mykey", Value: "myvalue", Path: "/abc/"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

//验证path属性是否生效的handler
func mypath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Cookies())
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	//路径必须以/abc/开头
	http.HandleFunc("/abc/mypath", mypath)
	server.ListenAndServe()
}

```

#### Expires

* Cookie默认存活时间是浏览器不关闭,当浏览器关闭后,Cookie失效
* 可以通过Expires设置具体什么时候过期,Cookie失效. 也可以通过MaxAge设置Cookie多长时间后实现
* IE6,7,8和很多浏览器不支持MaxAge,建议使用Expires
* Expires是time.Time类型,所以设置时需要明确设置过期时间
* 修改服务器端代码如下.只需要修改创建Cookie的代码,其他位置不变

```go
package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	//验证httponly
	//c := http.Cookie{Name: "mykey", Value: "myvalue", HttpOnly: false}
	//验证path
	//c := http.Cookie{Name: "mykey", Value: "myvalue", Path: "/abc/"}
	//验证Expires
	c := http.Cookie{Name: "mykey", Value: "myvalue", Expires: time.Date(2018, 1, 1, 1, 1, 1, 0, time.Local)}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

//验证path属性是否生效的handler
func mypath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Cookies())
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	//路径必须以/abc/开头
	http.HandleFunc("/abc/mypath", mypath)
	server.ListenAndServe()
}
```

