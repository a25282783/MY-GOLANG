package main

import (
	"fmt"
)

// 接口的名字由方法名加 [e]r or able后缀组成
type greeter interface {
	say(str string)
}

type american struct {
}

func (american *american) foo() {
	fmt.Println("american foo...")
}

func (american *american) say(str string) {
	fmt.Println("hi" + str)
}

type chinese struct {
}

func (chinese *chinese) bar() {
	fmt.Println("chinese bar...")
}

func (chinese *chinese) say(str string) {
	fmt.Println("hi" + str)
}

// 多態
func greet(g greeter, str string) {
	g.say(str)

	switch t := g.(type) {
	case *chinese:
		t.bar()
	case *american:
		t.foo()
	default:
	}
}

func main() {
	chinese := &chinese{}

	american := &american{}

	greet(chinese, "c")
	greet(american, "a")

	//=============
	var g greeter = chinese
	g.say("greeter")
	greet(g, "g")
}
