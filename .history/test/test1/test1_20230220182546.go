package main

import (
	"fmt"
)

func main() {
	a := Dog{
	}
	a.Shout()
}

type Animal struct {
	sound string
}

func (animal *Animal) Shout() string {
	return "..."
}

type Dog struct {
	sound = "woof"
	Animal
}

