package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/*
### 文件下载

* 文件下载总体步骤
	* 客户端向服务端发起请求，请求中包含要下载文件的名称
	* 服务端接收到请求后将文件设置到响应对象中，响应给客户端浏览器
* 下载时需要设置的响应头信息
  * Content-Type: 内容MIME类型
    * application/octet-stream 任意类型
  * Content-Disposition:客户端对内容的操作方式
    * inline 默认值,表示浏览器能解析就解析,不能解析下载
    * attachment;filename=下载时显示的文件名 ,客户端浏览器恒下载
*/

func showDownloadPage(writer http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(writer, nil)
}

func download(writer http.ResponseWriter, request *http.Request)  {
	// 获取请求参数
	filename := request.FormValue("filename")
	// 设置响应头
	header := writer.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("content-Disposition", "attachment;filename=" + filename)
	// 使用ioutil包读取文件
	b, err := ioutil.ReadFile("/Users/xiaoyuan/Desktop/testDir11/" + filename)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	// 将文件写入到响应对象中
	count, err := writer.Write(b)
	if err != nil {
		fmt.Println("写入文件失败", err)
	} else {
		fmt.Println(count)
	}
}

func main() {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", showDownloadPage)
	http.HandleFunc("/download", download)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}