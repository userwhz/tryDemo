package main

// gogenerate mockgen -destination=./mock_demo/mock_human.go -package=mock_demo -source=interface.go
type Human interface {
	Speak() string
	Walk() string
}
