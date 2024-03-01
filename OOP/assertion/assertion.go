package main

import "fmt"

// interface{} 万能数据类型
// 采用类型断言机制

func func1(arg interface{}) {
	fmt.Println("sss")
	fmt.Println(arg)
	value, ok := arg.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not string")
	}
}

type Book struct {
	title string
	edge  int
}

func main() {
	func1(111)
	book := Book{"name", 100}
	func1(book)
	fmt.Println()
	func1("hhh")
}
