package query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
数据库的查询
查询注意点：
Golang中执行查询与新增、删除、修改中stmt的执行方法有区别，由于需要把查询到的结果取出来，所以还需要进行取值处理
*/
func Demo()  {
	fmt.Println("开始查询数据库")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

	// 关闭数据库
	defer func() {
		if db != nil {
			db.Close()
		}
		fmt.Println("关闭数据库")
	}()

	err = db.Ping()
	if err != nil {
		fmt.Println("ping数据库失败", err)
		return
	}

	// 预处理，准备处理SQL语句，支持占位符，防止SQL注入
	stmt, err := db.Prepare("select * from people")
	if err != nil {
		fmt.Println("预处理失败", err)
		return
	}

	// 关闭stmt对象
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
		fmt.Println("关闭stmt对象")
	}()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("查询数据库失败", err)
		return
	}

	// 循环遍历结果
	for rows.Next() {
		var id int
		var name string
		var address string
		// 把行内值赋值给变量
		err := rows.Scan(&id, &name, &address)
		if err != nil {
			fmt.Println("查询结果集失败", err)
		} else {
			fmt.Println(id, name, address)
		}
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
		fmt.Println("关闭结果集")
	}()
}
