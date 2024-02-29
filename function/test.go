package main

import "fmt"

func fun1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

func func2(a int, b int) (r1 int, r2 int) {
	var result int
	result = a + b

	return result, a
}
func main() {
	a := 3
	b := 1
	c, d := func2(a, b)
	fmt.Println(c, d)
}
