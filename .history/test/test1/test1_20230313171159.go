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

	fmt.Println(Sum(Int32(1), 2))
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

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Sum[T Num](x, y T) T {
	return x + y
}

type Int32 = int64

func Map[T any](s []T, f func(t T) T) (r []T) {
	r = make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
