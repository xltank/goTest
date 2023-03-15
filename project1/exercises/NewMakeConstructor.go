package main

import "fmt"

type MyStruct struct {
	Name string
	Age  uint8
}

func NewMyStruct(name string, age uint8) MyStruct {
	if name == "" {
		name = "NoName"
	}
	if age < 0 {
		age = 0
	}
	return MyStruct{name, 0}
}

func main() {
	a := new(int)
	fmt.Println(a)

	b := new(MyStruct)
	fmt.Println(b)

	c := new([]int)
	fmt.Println(c)

	d := make([]int, 5, 10)
	fmt.Println(d)

	e := NewMyStruct("Jason", 33)
	fmt.Println(e)
	e.Age += 1
	fmt.Println(e)
}
