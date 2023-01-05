package main

import (
	"fmt"
	"my_golang/struct_demo/person"
)

func main() {
	s := person.NewStudent("A") // value
	t := person.NewTeacher("B") // ptr

	person.Test1(&s)
	person.Test2(t)

	fmt.Println(s)
	fmt.Println(t)
}
