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

func main() {
	int1 := 10
	foo(int1)

}
