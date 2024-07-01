package main

import (
	"github.com/agiledragon/gomonkey"
	"reflect"
	"testing"
)

type Boy struct {
}

func (b *Boy) Speak() string {
	return "aa"
}

func Test_gomonkey(t *testing.T) {
	b := Boy{}
	patch := gomonkey.ApplyMethod(reflect.TypeOf(&Boy{}), "Speak", func(b *Boy) string {
		return "5555"
	})

	defer patch.Reset()
	t.Errorf("speak: %s", b.Speak())
	t.Errorf("speak: %s", b.Speak())

}
