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

	fmt.Println(Add(Int64(1), 2, 3, 4))

	fmt.Println(Map([]string{"aa", "bb", "cc"}, func(s string) string {
		return s + "_" + s
	}))

	fmt.Println(Compact([]int{0, 1, 3, 5}))

	a2 := Animal{sound: "a"}
	fmt.Println(Animal{sound: "a"} == a2)

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
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Int64 int64

func Add[T Num](x ...T) T {
	var z T
	for i := 0; i < len(x); i++ {
		z += x[i]
	}
	return z
}

type Int32 = int64

func Map[T any](s []T, f func(t T) T) (r []T) {
	r = make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Compact[T any](s []T) (r []T) {
	r = make([]T, 0)
	// var zero T
	// for i, v := range s {
	// 	if zero == v {
	// 		r = append(r, s[i])
	// 	}
	// }
	fmt.Println(r)
	return r
}
