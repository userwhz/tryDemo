package main

import "fmt"

func main() {
	c := make(chan int)
	//1_go func() {
	//	for i := 0; i < 3; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
			close(c)
		}
	}()
	for {
		//ok为true说明channel没有关闭，ok为false表示channel关闭了
		//为了防止死锁，需要将channel关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main finish")

}
