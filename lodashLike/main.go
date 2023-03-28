package main

import (
	"fmt"
	"strconv"

	"lodashLike/slice"
)

var persons = []Person{
	{Name: "A", Age: 12},
	{Name: "AA", Age: 13},
	{Name: "B", Age: 18},
	{Name: "C", Age: 0},
	{Name: "B", Age: 18},
}

func main() {
	fmt.Println("Add()", slice.Add(Int64(1), 2, 3, 4))
	fmt.Println("Add()", slice.Add([]string{"a", "b", "c"}...))

	fmt.Println("Compact()", slice.Compact([]int{0, 1, 3, 5, 0.0}))
	fmt.Println("Compact()", slice.Compact([]string{"a", "b", "", "ccc", "阿萨德飞", ""}))

	fmt.Println("Difference()", slice.Difference([]string{"a", "b", "", "ccc", "阿萨德飞", ""}, []string{"a", "b"}))

	fmt.Println("FilterBy()", slice.FilterBy(persons, func(item Person) bool {
		return item.Name <= "B"
	}))

	fmt.Println("FindFirstIndex()", slice.FindFirstIndex(persons, Person{"B", 18}))
	fmt.Println("FindFirstIndexBy()", slice.FindFirstIndexBy(persons, func(item Person) bool {
		return item.Age == 18
	}))

	fmt.Println("Intersection()", slice.Intersection([]int{0, 1, 3, 5, 0.0}, []int{0, 2, 3, 6, 7, 0.0}))

	fmt.Println("Map()", slice.Map([]string{"aa", "bb", "cc"}, func(s string) string {
		return s + "_" + s
	}))
	fmt.Println("Map()", slice.Map([]int{1, 2, 3, 4, 5}, func(s int) string {
		return strconv.Itoa(s) + "_" + strconv.Itoa(s)
	}))

	fmt.Println("Reverse()", slice.Reverse([]string{"a", "b", "", "ccc", "阿萨德飞", ""}))

	fmt.Println("Sum()", slice.Sum(1, 2, 3, 4))

	fmt.Println("SumBy()", slice.SumBy(persons, "Age"))

	fmt.Println("Uniq()", slice.Uniq(persons))

	fmt.Println("Uniq()", slice.Uniq(persons))
}

type Int64 int64

type Person struct {
	Name string
	Age  int
}
