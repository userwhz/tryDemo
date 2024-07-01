package main

//go:generate mockgen -destination=./mock/mock_human.go -package=mock -source=interface.go
type Human interface {
	Speak() string
	Walk() string
}
