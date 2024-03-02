package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (this Student) Go() {
	fmt.Println("gogogo")
}

func reflectNum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func reflectClass(arg interface{}) {
	inputType := reflect.TypeOf(arg)
	fmt.Println(inputType)
	inputValue := reflect.ValueOf(arg)
	fmt.Println(inputValue)

	fmt.Println(111)
	//获取字段 先获取Field 再获取对应val
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name, field.Type)
		fmt.Println(value)
	}
	//获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var num float64 = 123.21
	reflectNum(num)
	stu := Student{"aa", 12}
	reflectClass(stu)
}
