package main

import (
	"fmt"
)

type Read interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read book")
}

func (this *Book) WriteBook() {
	fmt.Println("write book")
}

func main() {
	//{int 111}
	var aa int
	aa = 111
	var inf interface{}
	//{int 111}
	inf = aa
	value, _ := inf.(int)
	fmt.Println(value)

	// {Book }
	var r Read = &Book{}
	r.ReadBook()
	var w Write
	//{Book}断言成功是因为w和r的type是一样的
	w = r.(Write)
	w.WriteBook()
}
