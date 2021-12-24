package delete

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
mysql删除数据
删除注意点
- 删除和修改、新增结构一样，区别为SQL语句
- 在Go语言中要求如果要删除的数据不存在，那么RowsAffected()返回0
*/

func Demo()  {
	fmt.Println("开始执行删除语句")
	// 使用mysql驱动，连接名为first的数据库
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping数据库失败", err)
		return
	}

	// 关闭数据库
	defer func() {
		if db != nil {
			db.Close()
		}
		fmt.Println("数据库已关闭")
	}()

	/*
	预处理 准备SQL语句
	支持占位符，防止SQL注入
	*/
	stmt, err := db.Prepare("delete from people where id=?")
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

	// 执行预处理
	// EXec()参数为不定项，对应预处理中的占位符?个数
	res, err := stmt.Exec(1)
	if err != nil {
		fmt.Println("执行SQL语句失败", err)
		return
	}
	// 获取执行受影响的行数
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取执行结构失败", err)
		return
	}
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
