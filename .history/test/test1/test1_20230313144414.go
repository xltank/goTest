package main

import "fmt"

func main() {
	a := Dog{Animal: Animal{
		sound: "woof"},
	}
	fmt.Println(a.Shout())

	i := new(int)
	fmt.Println(*i)

	n := make(map[string]string)
	fmt.Println(n)
	n["a"] = "1"
	fmt.Println(n)

	m := new(map[string]string)
	fmt.Printf("%p", m)
	mm := *m
	fmt.Println(mm)
	mm["a"] = "aaa"
	fmt.Println(mm)
}

type Animal struct {
	sound string
}

func (animal *Animal) Shout() string {
	return animal.sound
}

type Dog struct {
	Animal
}
