package main

import "fmt"

func main() {
	a := Dog{Sound: "woof"}
	fmt.Println(a.Shout())
}

type Animal struct {
	Sound string
}

func (animal *Animal) Shout() string {
	return animal.Sound
}

type Dog struct {
	Sound string
	Animal
}
