package main

func main() {
	a := Dog{sound: "woof"}
	a.Shout()
}

type Animal struct {
	sound string
}

func (animal *Animal) Shout() string {
	return "..."
}

type Dog struct {
	sound string
	Animal
}
