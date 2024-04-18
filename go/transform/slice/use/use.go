package main

import (
	"fmt"
)

func main() {
	//切片的追加
	//slice := make([]int, 3, 5)
	//fmt.Println(len(slice), cap(slice), slice)
	//slice = append(slice, 12, 12)
	//fmt.Println(len(slice), cap(slice), slice)
	//slice = append(slice, 12, 12)
	//fmt.Println(len(slice), cap(slice), slice)
	//slice1 := make([]int, 3)
	//fmt.Println(len(slice1), cap(slice1), slice1)
	//slice1 = append(slice1, 1)
	//fmt.Println(len(slice1), cap(slice1), slice1)
	//fmt.Println()
	//切片的截取 左闭右开 s的改变不影响 s1 s1的改变影响s 浅拷贝
	s := []int{1, 2, 3, 4}
	s1 := s[0:2]
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println()
	//s = append(s, 1, 2, 3)
	//s = s[0:1]
	//s[0] = 100
	//fmt.Println(len(s), cap(s), s)
	//fmt.Println(len(s1), cap(s1), s1)
	s1[0] = 100
	s1 = append(s1, 1, 2, 3)
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(s1), cap(s1), s1)

	//相互影响
	//arr := [4]int{1, 2, 3, 4}
	//ss1 := arr[0:3]
	//ss2 := arr[0:4]
	//ss2[0] = 100
	//fmt.Println(arr[0], ss1[0], ss2[0])

	// slice容量小于256个元素，则扩容后容量直接翻倍，在roundupsize
	// slice容量不小于256个元素，则for每次增加原来容量是四分之一+四分之三的阈值，在roundupsize
	//ssK := make([]int, 1536)
	//fmt.Println(len(ssK), cap(ssK))
	//ssK = append(ssK, 2)
	//fmt.Println(len(ssK), cap(ssK))

	fmt.Println()

	//a := []int{1, 2, 3}
	//b := a
	//c := a[:]
	//a[0] = 100
	//fmt.Println(a, b, c) // [100 2 3] [100 2 3] [100 2 3]
	//深拷贝
	dd := []int{1, 2, 3}
	ddd := make([]int, 3)
	copy(ddd, dd)
	dd = append(dd, 2)
	ddd = append(ddd, 3)
	fmt.Println(ddd)
	fmt.Println(dd)
}
