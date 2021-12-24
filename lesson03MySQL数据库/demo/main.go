package main

import (
	"demo/src/insert"
	"demo/src/query"
	"demo/src/update"
	"demo/src/delete"
)

func main() {

	// 测试增加数据
	insert.Demo()
	update.Demo()
	delete.Demo()
	query.Demo()
}