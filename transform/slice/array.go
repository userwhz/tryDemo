package main

import "fmt"

func printArr(arr [10]int) {
	for index, value := range arr {
		fmt.Println("index = ", index, " ", value)
	}
}
func printSlice(arr []int) {
	for index, value := range arr {
		fmt.Println("slice = ", index, " ", value)
	}
	arr[0] = 100
}

func main() {
	//值传递
	var arr [10]int
	//slice
	arr1 := []int{1, 2, 3}

	//内置函数 不需要引入包
	print("aaa")
	fmt.Println()
	for index, value := range arr {
		fmt.Println(index, " ", value)
	}
	fmt.Printf("%T\n", arr)
	fmt.Println("bbb")
	printArr(arr)
	fmt.Printf("%T\n", arr1)
	//切片是引用传递
	printSlice(arr1)
	for index, value := range arr1 {
		fmt.Println("slice = ", index, " ", value)
	}
}
