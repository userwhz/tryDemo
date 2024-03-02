package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (this *Human) Eat() {
	fmt.Println("eat")
}

func (this *Human) Walk() {
	fmt.Println("walk")
}

type SuperMan struct {
	Human
	Level int
}

func (this *SuperMan) Eat() {
	fmt.Println("super eat")
}

func (this *SuperMan) Fly() {
	fmt.Println("super fly")
}
func (this *SuperMan) Print() {
	fmt.Println(this.name)
	fmt.Println(this.age)
	fmt.Println(this.Level)
}

func main() {
	human := Human{"ww", 12}
	human.Eat()
	human.Walk()

	superman := SuperMan{human, 1}

	superman.Walk()
	superman.Eat()
	superman.Fly()

	fmt.Println(human.name)
	fmt.Println(superman.name)

	human.name = "aa"
	fmt.Println(human.name)
	fmt.Println(superman.name)

	superman.Human.name = "bb"
	fmt.Println(human.name)
	fmt.Println(superman.name)

	var s SuperMan
	s.name = "ss"
	s.age = 1
	s.Level = 3

	s.Walk()
	s.Eat()
	s.Fly()
}
