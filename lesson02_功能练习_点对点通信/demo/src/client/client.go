package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type User struct {
	Username string
	OtherUsername string
	Msg string
	ServerMsg string
}

var (
	user = new(User)
	wg sync.WaitGroup
)

func main() {
	wg.Add(1)
	fmt.Println("请登录，输入用户名：")
	fmt.Scanln(&user.Username)
	fmt.Println("请输入要给谁发送消息")
	fmt.Scanln(&user.OtherUsername)
	addr, err := net.ResolveTCPAddr("tcp4", ":8889")
	if err != nil {
		fmt.Println("创建TCPAddress错误", err)
		return
	}
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println("建立连接失败", err)
		return
	}

	// 发送消息的子协程
	go func() {
		fmt.Println("请输入：（只提示一次，以后直接输入即可）")
		for {
			fmt.Scanln(&user.Msg)
			if user.Msg == "exit" {
				// 退出
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			// 发送消息
			conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
		}
	}()

	// 接收消息的子协程
	go func() {
		for {
			b := make([]byte, 512)
			count, err := conn.Read(b)
			if err != nil {
				fmt.Println("读取消息错误", err)
				continue
			}
			user2 := new(User)
			arr := strings.Split(string(b[:count]), "-")
			user2.Username = arr[0]
			user2.OtherUsername = arr[1]
			user2.Msg = arr[2]
			user2.ServerMsg = arr[3]
			if user2.ServerMsg != "" {
				fmt.Println("\t\t服务器消息:", user2.ServerMsg)
			} else {
				fmt.Println("\t\t", user2.Username, user2.Msg)
			}
		}
	}()

	wg.Wait()
}