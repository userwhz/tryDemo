package main

import "fmt"

func main() {
	var map1 map[string]string
	if map1 == nil {
		fmt.Println(map1)
	}
	map1 = make(map[string]string, 10)
	map1["one"] = "1"
	map1["two"] = "2"
	map1["three"] = "3"
	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["a"] = 1
	map2["b"] = 2
	fmt.Println(map2)

	map3 := map[string]int{
		"1": 1,
		"2": 2,
	}
	fmt.Println(map3)

	fmt.Println()
	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"
	val, key := language["php"] //查找是否有php这个子元素
	fmt.Println(val, key)
	if key {
		fmt.Printf("%v", val)
	} else {
		fmt.Printf("no")
	}
}
