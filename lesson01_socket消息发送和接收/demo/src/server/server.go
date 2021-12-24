package main

/*
服务端代码
*/

import (
	"fmt"
	"net"
)

func main()  {
	// 创建TCPAddress变量，指定协议tcp4，监听本机8899
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8899")
	if err != nil {
		fmt.Println("创建TCPAddress失败:", err)
		return
	}

	// 监听TCPAddress设定的地址
	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("监听TCPAddress地址失败:", err)
		return
	}

	fmt.Println("服务器已经启动")

	// 阻塞式等待客户端消息，返回连接对象，用于接收客户端消息或箱客户端发送消息
	conn, err := listener.Accept()
	defer func() {
		// 关闭连接
		_ = conn.Close()
	}()
	if err != nil {
		fmt.Println("接收数据错误：", err)
		return
	}

	// 把数据读取到切片中
	b := make([]byte, 256)
	fmt.Println("read之前")

	// 客户端没有发送数据且客户端对象没有关闭，Read()将会阻塞，一旦接收到数据就不阻塞
	count, err := conn.Read(b) // 返回读取数据的大小
	if err != nil {
		fmt.Println("读取数据错误：", err)
		return
	}
	fmt.Println("接收到的数据：", string(b[:count]))

	fmt.Println("服务结束")
}