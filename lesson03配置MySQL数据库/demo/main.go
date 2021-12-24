package main

import (
	"database/sql"
	"fmt"
	// 驱动已经放入到标准库文件夹,由于不使用所以需要空导入, 在前面添加`_`
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	insertDemo()
}


/*
演示，
连接本地localhost:3306下的名为`first`的数据库
向people中新增一条数据
*/
func insertDemo()  {
	// 1.打开数据库
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("错误", err)
	}
	// 关闭连接
	defer func() {
		if db != nil {
			_ = db.Close()
		}
		fmt.Println("关闭连接")
	}()

	/*
		准备处理SQL语句
		支持占位符，防止SQL注入
	*/
	stmt, err := db.Prepare("insert into people values(default , ?, ?)")
	if err != nil {
		fmt.Println("预处理失败", err)
		return
	}
	// 关闭对象
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()

	/*
		Exec()函数的参数为不定项参数，对应占位符?个数
	*/
	res, err := stmt.Exec("张三", "海淀")
	if err != nil {
		fmt.Println("执行sql语句出现错误", err)
		return
	}

	// 获取新增时生成的主键的值
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("获取主键失败", err)
		return
	}

	// 受影响的行数
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败", err)
		return
	}
	fmt.Println("受影响的行数:", count, "主键为:", id)
	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}
}