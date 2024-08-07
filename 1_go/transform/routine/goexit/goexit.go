package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a, b int) {
		fmt.Println(a, b)
	}(10, 2)

	func() {
		time.Sleep(1 * time.Second)
	}()

}
