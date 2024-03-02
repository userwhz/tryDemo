package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听msg广播消息的goroutine
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)

	//用户上线
	user.Online()

	isLive := make(chan bool)

	// 接受客户端发送的消息并进行全体广播
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//用户下线
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err", err)
				return
			}

			msg := string(buf[:n-1])
			//用户对msg进行处理
			user.DoMessage(msg)

			//存在消息 说明用户是活跃的
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
			//do nothing
		case <-time.After(100 * time.Second):
			// 已经超时
			user.SendMsg("you are out")
			close(user.C)
			conn.Close()
			return
		}
	}

}

// 启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()
	//启动msg的监听
	go this.ListenMessager()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err", err)
			continue
		}
		//do handler
		go this.Handler(conn)

	}
	// defer 处理 close listen socket
}
