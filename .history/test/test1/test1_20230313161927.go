package main

import "fmt"

func main() {
	a := Dog{Animal: Animal{
		sound: "woof"},
	}
	fmt.Println(a.Shout())

	t := true
	f := false
	fmt.Println(t == f)

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