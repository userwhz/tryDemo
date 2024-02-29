package main

import (
	"fmt"
	"tryDemo/init/lib1"
	//mylib2 "tryDemo/init/lib2"
	aaa "tryDemo/init/lib2"
)

func changeVal(a *int) {
	*a = 100
}
func main() {
	lib1111.Lib1Test()
	aaa.Lib2Test()
	aaa.Lib2Test()
	var a = 10
	changeVal(&a)
	fmt.Print(a)
}

//参数传递，默认使用的是值传递
