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

	fmt.Println(Add("aa", "bb", "cc"))

	fmt.Println(Map([]string{"aa", "bb", "cc"}, func(s string) string {
		return s + "_" + s
	}))

	fmt.Println(Compact([]int{0, 1, 3, 5}))

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

type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func Add[T Addable](x ...T) T {
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
	for i, v := range s {
		switch d := any(v).(type) {
		case string:
			if d == "" {
				continue
			}
		case int:
			if d == 0 {
				continue
			}
		case int32:
			if d == 0 {
				continue
			}
		case int64:
			if d == 0 {
				continue
			}
		case float32:
			if d == 0 {
				continue
			}
		case float64:
			if d == 0 {
				continue
			}
		}
		r = append(r, s[i])
	}
	fmt.Println(r)
	return r
}
