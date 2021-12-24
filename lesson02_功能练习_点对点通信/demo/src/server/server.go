package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Username string
	OtherUsername string
	Msg string
	ServerMsg string
}

var (
	userMap = make(map[string]net.Conn)
	user = new(User)
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8889")
	if err != nil {
		fmt.Println("创建TCPAddress失败")
		return
	}
	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("监听TCPAddress失败", err)
		return
	}

	for {
		fmt.Println("服务器接收消息")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接收消息失败", err)
			continue
		}
		fmt.Println("服务器接收到消息")
		go func() {
			for {
				b := make([]byte, 512)
				// 读取数据
				count, err := conn.Read(b)
				if err != nil {
					fmt.Println("读取消息失败", err)
					continue
				}
				arr := strings.Split(string(b[:count]), "-")
				user.Username = arr[0]
				user.OtherUsername = arr[1]
				user.Msg = arr[2]
				user.ServerMsg = arr[3]
				userMap[user.Username] = conn
				if v, ok := userMap[user.OtherUsername]; ok && v != nil {
					// 对方在线时，将用户A发送到服务器的消息，发送给用户B
					user.ServerMsg = ""
					c, e := v.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
					if c == 0 || e != nil {
						// 发送失败时，关闭连接
						_ = conn.Close()
						delete(userMap, user.OtherUsername)
						break
					}

				} else {
					// 对方不在线时，告诉用户A对方不在线
					user.ServerMsg = "对方不在线"
					c, e := conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
					if c == 0 || e != nil {
						fmt.Println("发生错误", e)
					}
				}
			}
		}()
	}

	fmt.Println("服务器结束")
}
