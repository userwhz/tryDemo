package main

import "fmt"

const (
	//iota 每行都会累加1
	A    = 10 * iota
	B    = "asd"
	C    = iota
	D    = iota
	E, F = iota + 1, iota
	G    = iota
)

func main() {
	const length int = 10
	fmt.Println(length)
	fmt.Println(A, B, C, D)
}
