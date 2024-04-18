package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		defer fmt.Println("发送完成")
		for i := 0; i < 5; i++ {
			fmt.Println("正在发送")
			c <- i
			fmt.Println("数据", i, "长度", len(c), "容量", cap(c))
		}
	}()
	time.Sleep(3 * time.Second)
	func() {
		for i := 0; i < 5; i++ {
			num := <-c
			fmt.Println("接受", num)
		}
	}()

	fmt.Println("main over")

}
