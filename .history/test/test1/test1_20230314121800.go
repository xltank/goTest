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

	fmt.Println(Map([]string{"aa", "bb", "cc"}, func(s string) string {
		return s + "_" + s
	}))

	fmt.Println(Compact([]int32{0, 1, 3, 5, 0}))

	var i Int32 = 10
	fmt.Printf("%T\n", i)
	var j Int64 = 10
	fmt.Printf("%T\n", j)
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

type Int32 = int32
type Int64 int64

func Map[T any](s []T, f func(t T) T) (r []T) {
	r = make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Compact[T Num](s []T) (r []T) {
	r = make([]T, 0)
	for i, v := range s {
		if v == 0 {
			continue
		}
		r = append(r, s[i])
		fmt.Println(r)
	}
	return r
}
