package main

import "fmt"

func func1() int {

	defer fmt.Println("111")
	return 0
}

func func2() int {
	fmt.Println("222")
	return 0
}
func func3() int {
	defer func1()
	defer func4()
	return func2()
}
func func4() int {
	fmt.Println("333")
	return 0
}
func reco(i int) {
	var arr [10]int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 10
	arr[i+1] = 11

}
func main() {
	//defer压栈进入
	defer fmt.Println("end1")
	defer fmt.Println("end2")

	fmt.Println("start")
	defer fmt.Println("end3")
	func3()
	reco(11)
	fmt.Println("继续执行")
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}
