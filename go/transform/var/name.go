//四种变量的声明方式

package main

import "fmt"

func main() {
	var a int
	fmt.Println("a =", a)
	fmt.Printf("%T\n", a)

	var b int = 10
	fmt.Println("b =", b)
	fmt.Printf("%T\n", b)

	var c = "cc"
	fmt.Println("c =", c)
	fmt.Printf("%T\n", c)

	var d string = "ads"
	fmt.Println("d =", d)
	fmt.Printf("%T\n", d)

	//省去var关键字，直接自动匹配 不能声明全局变量
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("%T\n", e)

	f := 3.44
	fmt.Println("f =", f)
	fmt.Printf("%T\n", f)

	var aa, bb int = 100, 200
	fmt.Println(aa)
	fmt.Println(bb)
	var cc, dd = 123, "dd"
	fmt.Println(cc)
	fmt.Println(dd)
	var (
		vv int  = 100
		jj bool = true
	)

	fmt.Println(vv)
	fmt.Println(jj)
}
