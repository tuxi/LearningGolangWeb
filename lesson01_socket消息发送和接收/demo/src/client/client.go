package main

import (
	"fmt"
	"net"
	"time"
)

/*
客户端代码
*/

func main()  {

	// 服务器端id和端口
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8899")
	if err != nil {
		fmt.Println("创建失败", err)
		return
	}

	// 申请连接服务端
	coon, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println("申请连接服务器失败", err)
		return
	}

	// 向服务端发送数据
	count, err := coon.Write([]byte("Hello world，我的第一个goweb程序"))
	if err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	fmt.Println("客户端向服务端发送的数据量为：", count)
	/*
	通过休眠测试客户端对象不关闭，服务器是否能接收到对象
	测试结果，不影响服务端接收
	 */
	time.Sleep(10 * time.Second)

	// 关闭连接
	_ = coon.Close()
}