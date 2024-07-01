package main

import (
	"context"
	"fmt"
	"net"
	greeter "tryDemo/7_RPC/gprc/proto"

	"google.golang.org/grpc"
)

// 声明GreeterImpl结构体，并实现SayHello方法
type GreeterImpl struct {
	// 将greeterUnimplementedGreeterServer内置
	greeter.UnimplementedGreeterServer
}

// 实现SayHello方法
func (g *GreeterImpl) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	fmt.Println("req:", req.Data)
	return &greeter.HelloReply{Data: "hello， client"}, nil
}

func main() {
	// 创建listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 创建grpc server
	server := grpc.NewServer()

	// 注册Greeter服务到server中
	greeter.RegisterGreeterServer(server, &GreeterImpl{})

	// 启动服务
	if err = server.Serve(listener); err != nil {
		fmt.Println("Server:", err.Error())
	}
}
