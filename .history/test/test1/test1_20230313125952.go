package main

import "fmt"

func main() {
	a := Dog{Animal: Animal{
		sound: "woof"},
	}
	fmt.Println(a.Shout())

	i := new(int)
	fmt.Println(*i)

	m := new(map[string]string)
	mm := *m
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
