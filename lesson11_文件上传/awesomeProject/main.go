package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/*
# 文件上传

* 文件上传:客户端把上传文件转换为二进制流后发送给服务器,服务器对二进制流进行解析
* HTML表单(form)enctype(Encode Type)属性控制表单在提交数据到服务器时数据的编码类型.
  * enctype=”application/x-www-form-urlencoded” 默认值,表单数据会被编码为名称/值形式
  * enctype=”multipart/form-data” 编码成消息,每个控件对应消息的一部分.请求方式必须是post
  * enctype=”text/plain” 纯文本形式进行编码的
*/

func welcome(writer http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(writer, nil)
}

func upload(writer http.ResponseWriter, request *http.Request)  {
	// 获取普通表单的数据
	username := request.FormValue("username")
	password := request.FormValue("password")
	fmt.Println(username, password)
	// 读取文件流
	avatar, header, _ := request.FormFile("avatar")
	// 将文件流转换为字节切片
	b, _ := ioutil.ReadAll(avatar)
	// 把文件保存到指定位置
	newPath := "/Users/xiaoyuan/Desktop/testDir11/" + header.Filename
	err := ioutil.WriteFile(newPath, b, 0777)
	if err != nil {
		fmt.Println("文件上传失败", err)
		t, _ := template.ParseFiles("view/failure.html")
		t.Execute(writer, err)
	} else {
		// 输出上传时的文件名称
		fmt.Println("文件上传成功", header.Filename)
		t, _ := template.ParseFiles("view/success.html")
		t.Execute(writer, newPath)
	}


}

func main() {
	server := http.Server{
		Addr: ":8090",
	}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/upload", upload)
	server.ListenAndServe()
}