package iInterface

import (
	"fmt"
)

func Test() {
	fmt.Println("interface demo begin: \n", new(World), new(World).SayHi(), printHi())
}

type World struct {
	info string
}

type helloWorld interface {
	SayHi() string
}

// func (w *World) String() string {
// 	return "hello,世界"
// }

func (w *World) SayHi() string {
	return "hi"
}

func printHi() string {
	para := World{
		info: "I'm the World",
	}
	return getHi(&para)
}

func getHi(param helloWorld) string {

	return fmt.Sprintf("%+v,%s\n", param, param.SayHi())
}
