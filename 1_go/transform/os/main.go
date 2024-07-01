package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个信号通道
	sigChan := make(chan os.Signal, 1)

	// 注册要监听的信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个 goroutine 以等待信号
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %s\n", sig)
		// 处理信号（可以在这里进行清理操作等）
		os.Exit(0)
	}()

	// 模拟程序运行
	fmt.Println("Program is running. Press Ctrl+C to exit.")
	select {}
}
