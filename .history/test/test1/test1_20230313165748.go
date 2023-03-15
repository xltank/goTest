package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/condition"
)

func main() {
	a := Dog{Animal: Animal{
		sound: "woof"},
	}
	fmt.Println(a.Shout())

	t := true
	f := false
	fmt.Println(t == f)

	fmt.Println("TernaryOperator:", condition.TernaryOperator(1 > 0, "true", "false"))
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

type Living interface {
	~Animal
}
