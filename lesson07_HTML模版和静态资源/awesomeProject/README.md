### 项目结构

* 在Go语言中web项目标准结构如下
```
--项目名称
    --src
    --static
        --css
        --js
    --view
        --index.html
    --main.go
```
* Go语言标准库中html/template包提供了html模版支持，把html当作模版可以在访问控制器时显示HTML模版信息
    * 这也符合标准的MVC思想
    
### HTML模版显示

* 使用template.ParseFiles()可以解析多个模版文件
```go
// ParseFiles creates a new Template and parses the template definitions from
// the named files. The returned template's name will have the (base) name and
// (parsed) contents of the first file. There must be at least one file.
// If an error occurs, parsing stops and the returned *Template is nil.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
// named "foo", while "a/foo" is unavailable.
func ParseFiles(filenames ...string) (*Template, error) {
	return parseFiles(nil, filenames...)
}
```
* 把模版信息响应写入到输出流中
```go
// Execute applies a parsed template to the specified data object,
// writing the output to wr.
// If an error occurs executing the template or writing its output,
// execution stops, but partial results may already have been written to
// the output writer.
// A template may be executed safely in parallel, although if parallel
// executions share a Writer the output may be interleaved.
func (t *Template) Execute(wr io.Writer, data interface{}) error {
	if err := t.escape(); err != nil {
		return err
	}
	return t.text.Execute(wr, data)
}
```

### 引用静态文件
* 把静态文件放入到特定的文件夹中,使用Go语言的文件服务就可以进行加载
* 项目结构
```
--项目
	--static
		--js
			--index.js
	--view
		--index.html
	--main.go
```
* index.html代码如下
```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
    <!--路径以斜杠开头,表示项目根目录下-->
    <script type="text/javascript" src="/static/js/index.js"></script>
</head>
<body>
    这是要显示的html页面信息<br/>
    <button onclick="myclick()">按钮</button>
</body>
</html>
```
* index.js代码如下
```javascript
function myclick(){
    alert("您点击了按钮")
}
```
    
