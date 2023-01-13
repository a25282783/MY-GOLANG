package main

import (
	"fmt"
	"reflect"
)

type any interface {
}

func foo(any any /*任何東西*/) {
	refT := reflect.TypeOf(any)
	refV := reflect.ValueOf(any)

	fmt.Println(refT.Name()) // int
	fmt.Println(refV)        // 10
	// fmt.Println(refV + 2)    // wrong
	// fmt.Println(refV.Interface() + 2) // wrong
	fmt.Println(refV.Interface().(int) + 2) // 12 轉成空接口再斷言
}

func getKind(any interface{} /*簡寫*/) {
	refT := reflect.TypeOf(any)

	fmt.Println(refT.Kind())
}

func main() {
	// int1 := 10
	// foo(int1)
	// getKind(int1)

	type A struct {
	}
	a := &A{}
	getKind(*a) // struct
	b := &A{}
	getKind(b) // ptr

	slice1 := []int{0, 1}
	getKind(slice1) // slice

}
