package main

import (
	"fmt"
	"strconv"

	"lodashLike/slice"
)

func main() {
	fmt.Println("Add()", slice.Add(Int64(1), 2, 3, 4))
	fmt.Println("Add()", slice.Add("a", "b", "c"))

	fmt.Println("Sum()", slice.Sum(1, 2, 3, 4))

	fmt.Println("Map()", slice.Map([]string{"aa", "bb", "cc"}, func(s string) string {
		return s + "_" + s
	}))
	fmt.Println("Map()", slice.Map([]int{1, 2, 3, 4, 5}, func(s int) string {
		return strconv.Itoa(s) + "_" + strconv.Itoa(s)
	}))

	fmt.Println("Compact()", slice.Compact([]int{0, 1, 3, 5, 0.0}))
	fmt.Println("Compact()", slice.Compact([]string{"a", "b", "", "ccc", "阿萨德飞", ""}))

	fmt.Println("Compact2()", slice.Compact2([]int{0, 1, 3, 5, 0.0}))

	persons := []Person{
		{Name: "A", Age: 12.1},
		{Name: "B", Age: 18.2},
		{Name: "C", Age: 18.11111},
	}
	fmt.Println("SumBy()", slice.SumBy(persons, "Age"))
}

type Int64 int64

type Person struct {
	Name string
	Age  float32
}
