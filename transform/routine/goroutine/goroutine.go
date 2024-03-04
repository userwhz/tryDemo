package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(2 * time.Second)
	}
}
