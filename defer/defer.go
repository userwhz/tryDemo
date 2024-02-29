package main

import "fmt"

func main() {
	//defer压栈进入
	defer fmt.Println("end1")
	defer fmt.Println("end2")

	fmt.Println("start")
	defer fmt.Println("end3")
}
