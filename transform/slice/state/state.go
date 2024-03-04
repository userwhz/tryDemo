package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Println(len(slice1), slice1)
	var slice2 []int
	fmt.Println(len(slice2), slice2)
	slice2 = make([]int, 3)
	fmt.Println(len(slice2), slice2)
	slice3 := make([]int, 3)
	fmt.Println(len(slice3), slice3)
	var slice4 []int
	if slice4 == nil {
		fmt.Println("空切片")
	} else {
		for index, value := range slice3 {
			fmt.Println(index, value)
		}
	}
}
