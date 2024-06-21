package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `inf:"name" doc:"名字"`
	Sex  string `inf:"boy"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		inf := t.Field(i).Tag.Get("inf")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println(inf, doc)
	}
}

func main() {
	res := Resume{"ww", "boy"}
	findTag(&res)
}
