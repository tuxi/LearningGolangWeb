package update

import (
	"database/sql"
	"fmt"
)

/*
演示修改people表中的一条数据
将people表中id为1的数据修改
*/

func Demo()  {
	fmt.Println("开始执行修改语句")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

  	err = db.Ping()
	if err != nil {
		fmt.Println("数据库ping失败", err)
		return
	}

	// 关闭连接
	defer func() {
		if db != nil {
			_ = db.Close()
		}
		fmt.Println("数据库关闭")
	}()

  	/*
  	准备处理SQL语句
  	支持占位符，防止SQL注入
  	*/
  	stmt, err := db.Prepare("update people set name=?, address=? where id=?")
	if err != nil {
		fmt.Println("预处理失败", err)
		return
	}
	// 关闭对象stmt
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
		fmt.Println("关闭stmt")
	}()

  	// 执行语句失败
  	res, err := stmt.Exec("ivanns", "北京市房山区", 1)
	if err != nil {
		fmt.Println("执行SQL失败")
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("查询受影响的行数失败", err)
		return
	}
	fmt.Println("受影响的行数：", count)
	if count > 0 {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}
}