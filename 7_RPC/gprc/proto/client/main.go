package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	greeter "tryDemo/6_rpc/gprc/proto"
)

func main() {
	// 1. 使用Dial连接服务端
	conn, err := grpc.Dial("127.0.0.1:8080",
		// 1.1 创建不安全的证书
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 1.2 阻塞直到连接建立成功
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建出greeter Service
	greeterClient := greeter.NewGreeterClient(conn)

	// 调用SayHello方法
	reply, err := greeterClient.SayHello(context.Background(), &greeter.HelloRequest{Data: "hello, server"})
	if err != nil {
		slog.Error("call SayHello", "error", err)
		return
	}

	fmt.Println("server reply ", reply.Data)

}
