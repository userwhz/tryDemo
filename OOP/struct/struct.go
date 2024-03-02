package main

import "fmt"

type myint int

// 大写其他包能访问，小写只有包内访问
type Book struct {
	title  string
	author string
}

// 传递一个book的副本
func changeBook(book Book) {
	book.title = "Go"
}

func changeBook2(book *Book) {
	book.title = "Go"
}
func main() {
	var a myint = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var book1 Book
	book1.title = "Name"
	book1.author = "www"
	fmt.Println(book1)
	changeBook(book1)
	fmt.Println(book1)
	changeBook2(&book1)
	fmt.Println(book1)
}
