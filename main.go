package main

import (
	"fmt"
	"my_golang/struct_demo/person"
)

func main() {
	// sum, sub := lib1.Calc(1, 2)
	// println(sum, sub)

	// if res, err := lib1.TestErr(); err == nil {
	// 	fmt.Println(res)
	// } else {
	// 	fmt.Println(err)
	// }

	// if res := lib1.GetInt(); res != nil {
	// 	fmt.Println("GetInt Wrong")
	// } else {
	// 	fmt.Println(res)
	// }

	// map1 := lib1.MakeMap()
	// map2 := map1
	// map2[0]["name"] = "shit"
	// for k, p := range map1 {
	// 	fmt.Printf("%v: i am %v, age %v \n", k, p["name"], p["age"])
	// }

	// 淺拷貝
	// p1 := lib1.MakeTeacher(18, "Tim")
	// p2 := p1
	// p2.Age = 50
	// p2.Attr["move"] = "left"
	// fmt.Printf("p1: %v \n", p1)
	// fmt.Printf("p2: %v \n", p2)
	// fmt.Println(p1.Run())
	// 深拷貝
	// p1 := lib1.MakeTeacher(18, "Tim")
	// p2 := p1
	// p2.Age = 50
	// p2.Attr["move"] = "left"
	// fmt.Printf("p1: %v \n", p1)
	// fmt.Printf("p2: %v \n", p2)

	s := person.NewStudent("A")
	t := person.NewTeacher("B")

	s1 := s
	s1.Name = "C"

	fmt.Println(s)
	fmt.Println(t)

}
