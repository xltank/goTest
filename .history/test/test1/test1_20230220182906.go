package main

import "fmt"

func main() {
	a := Dog{sound: "woof"}
	fmt.Println(a.Shout())
}

type Animal struct {
	sound string
}

func (animal *Animal) Shout() string {
	return animal.sound
}

type Dog struct {
	sound string
	Animal
}
