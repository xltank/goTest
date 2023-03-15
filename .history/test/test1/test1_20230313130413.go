package main

import "fmt"

func main() {
	a := Dog{Animal: Animal{
		sound: "woof"},
	}
	fmt.Println(a.Shout())

	i := new(int)
	fmt.Println(*i)

	fmt.Println(make(map[string]string))

	m := new(map[string]string)
	fmt.Println(m)
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
