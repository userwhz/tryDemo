package main

import (
	"fmt"
	"sync"
)

func main() {
	var sMp sync.Map
	sMp.Store("k1", "v1")
	sMp.Delete("k1")
	v, ok := sMp.Load("k1")
	if !ok {
		fmt.Println("k1 not exist")
	}
	fmt.Println(v)
	sMp.Store("k1", "v1")
	sMp.Store("k2", "v2")
	sMp.Store("k3", "v3")
	sMp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return key != "k2"
	})

}
