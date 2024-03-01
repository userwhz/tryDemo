package main

import "fmt"

func main() {
	//内部具有同步机制
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine over")
		c <- 666
	}()
	num := <-c
	fmt.Println(num)
	fmt.Println("main over")
}
