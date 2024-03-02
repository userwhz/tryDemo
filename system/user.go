package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListemMessage()
	return user
}

// 用户的上线业务
func (this *User) Online() {
	//用户上线，将用户加入map
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播用户上线消息
	this.server.BroadCast(this, "online")
}

// 用户的下线业务
func (this *User) Offline() {
	//用户下线，将用户从map中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播用户下线消息
	this.server.BroadCast(this, "offline")
}

// 给当前客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线人物
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online..\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式:rename|zhangsan
		newName := strings.Split(msg, "|")[1]
		//判断name是否唯一
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("already exist\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("success update" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式 to|zhangsan|msg
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("Incorrect format, like to|zhangsan|...\n")
			return
		}
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("user is not exist")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("invalid information\n")
			return
		}
		remoteUser.SendMsg(this.Name + " say to you:" + content)
	} else {
		this.server.BroadCast(this, msg)
	}

}
func (this *User) ListemMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
