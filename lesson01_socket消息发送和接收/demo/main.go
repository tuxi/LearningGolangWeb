package main

/*
# 一.Socket简介

* 在标准库的net包中可供了可移植的网络I/O接口,其中就包含了Socket
* Socket在TCP/IP网络分层中并不存在,是对TCP或UDP封装
* 如果非要给Socket一个解释
  * 实现网络上双向通讯连接的一套API
  * 常称Socket为"套接字"
* Socket分类:
  * 按照连接时间
    * 短连接
    * 长连接(HTTP 1.1开始也支持长连接,Socket替换方案)
  * 按照客户端和服务器端数量
    * 点对点
    * 点对多
    * 多对多
* 网络通信内容都是包含客户端和服务端,服务端运行在服务器中,而客户端运行在客户端中,可以是浏览器,可以是桌面程序,也可以是手机App.客户端和服务端进行数据交互遵守特定的协议.

# 二.Go语言对Socket的支持

* TCPAddr结构体表示服务器IP和端口
  * IP是`type IP []byte`
  * Port是服务器监听的接口
```go
// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}
```
* TCPConn结构体表示连接,封装了数据读写操作
```go
// TCPConn is an implementation of the Conn interface for TCP network
// connections.
type TCPConn struct {
	conn
}
```
* TCPListener负责监听服务器端特定端口
```go
// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
	fd *netFD
}
```
*/

func main()  {

}