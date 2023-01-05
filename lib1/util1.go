package lib1

import (
	"errors"
)

type Person struct {
	Name string
	Age  int
	Attr map[string]string
}

func Calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func TestErr() (int, error) {
	rs := 0

	if rs < 1 {
		err := errors.New("TestErr wrong!")
		return 0, err
	}

	return rs, nil
}

func GetInt() int {
	defer handlePanic()

	a := 10
	b := 0
	c := a / b

	return c
}

func handlePanic() int {
	if a := recover(); a != nil {
		return -1
	}

	return -1
}
func MakeMap() map[int]map[string]string {
	map1 := make(map[int]map[string]string)

	map1[0] = map[string]string{
		"name": "Joe",
		"age":  "20",
	}

	map1[1] = map[string]string{
		"name": "Tim",
		"age":  "22",
	}

	return map1
}

func MakeTeacher(age int, name string) Person {
	p := Person{
		Name: name,
		Age:  age,
		Attr: map[string]string{
			"move": "right",
		},
	}
	return p
}

func MakePtrTeacher(age int, name string) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	p.Attr = map[string]string{
		"move": "right",
	}
	return p
}

func (p Person) Run() string {
	return p.Name + "running.."
}

func (p *Person) ChangePersonName() {
	p.Name = "changePersonName"
}
