package main

import (
	"fmt"
)

func fibonacci(c chan int, quit chan int) {
	x, y, z := 1, 1, 1
	for {
		select {
		case <-c:
			//x, y = y, x+y
			z = y
			y = x + y
			x = z
			fmt.Println("y", y)
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- 1
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	fmt.Println()
	x := 1
	y := 2
	xx := 1
	yy := 2
	xx = yy
	yy = xx + yy

	fmt.Println(x, y)
	fmt.Println(xx, yy)

}
