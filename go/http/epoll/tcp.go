package main

import "net"

// 启动一个 tcp 服务端代码示例
func main() {
	// 创建一个 tcp 端口监听器
	l, _ := net.Listen("tcp", ":8080")
	// 主动轮询模型
	for {
		// 等待 tcp 连接到达
		conn, _ := l.Accept()
		// 开启一个 goroutine 负责一笔客户端请求的处理
		go serve(conn)
	}
}

// 处理一笔 tcp 连接
func serve(conn net.Conn) {
	defer conn.Close()
	var buf []byte
	// 读取连接中的数据
	_, _ = conn.Read(buf)
	// ...
}
