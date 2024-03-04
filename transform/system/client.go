package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func (this *Client) SelecUsers() {
	sendMsg := "who\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err", err)
		return
	}
}

func (this *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	this.SelecUsers()
	fmt.Println("input username, if you want to exit, input exit")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println("input msg, if you want to exit, input exit")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := this.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write err", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("input msg, if you want to exit, input exit")
			fmt.Scanln(&chatMsg)
		}
		this.SelecUsers()
		fmt.Println("input username, if you want to exit, input exit")
		fmt.Scanln(&remoteName)
	}

}
func (this *Client) PublicChat() {
	var chatMsg string
	fmt.Println("input msg, if you want to out, input exit")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {

		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := this.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("input msg, if you want to out, input exit")
		fmt.Scanln(&chatMsg)
	}

}
func (this *Client) DealReaponse() {
	//一旦有数据，就拷贝到标准输出上，永久阻塞监听
	io.Copy(os.Stdout, this.conn)
}
func (this *Client) UpdateName() bool {
	fmt.Println("input username:")
	fmt.Scanln(&this.Name)
	sendMsg := "rename|" + this.Name + "\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err", err)
		return false
	}
	return true
}
func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {

		}
		switch this.flag {
		case 1:
			this.PublicChat()
			break
		case 2:
			this.PrivateChat()
			break
		case 3:
			this.UpdateName()
			break
		}
	}
}
func (this *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println("invalid input")
		return false
	}
}
func NewClinet(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       10086,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err", err)
		return nil
	}
	client.conn = conn
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server Ip (default 127.0.0.1) ")
	flag.IntVar(&serverPort, "port", 8888, "set server Port (default 8888) ")
}
func main() {
	client := NewClinet(serverIp, serverPort)
	if client == nil {
		fmt.Println("connect to server fail")
	} else {
		fmt.Println("connect to server success")
	}
	go client.DealReaponse()
	client.Run()

}
