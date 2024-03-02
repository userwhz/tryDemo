package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetName() string
	GetType()
}

type Cat struct {
	name string
}

func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}

func (this *Cat) GetName() string {
	return this.name
}

func (this *Cat) GetType() {
	fmt.Println("cat")
}

type Dog struct {
	name string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog sleep")
}

func (this *Dog) GetName() string {
	return this.name
}

func (this *Dog) GetType() {
	fmt.Println("Dog")
}

func main() {
	var animal AnimalIF
	cat := Cat{"cc"}
	animal = &cat
	animal.Sleep()
	animal.GetType()
	fmt.Println(animal.GetName())

	animal = &Dog{"ddd"}
	animal.Sleep()
	animal.GetType()
	fmt.Println(animal.GetName())
}
