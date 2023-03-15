// main
package main

import (
	"fmt"
)

func main() {

	print()
	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[0:3]
	s = append(s, 6)
	fmt.Println("s:", s)
	fmt.Println("arr:", arr)

	t := arr[1:]
	fmt.Println("t:", t)
	t = append(t, 7)
	fmt.Println("t:", t)
	fmt.Println("arr:", arr)

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3, 10)
	copy(b, a)
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b))
	c := append(a, b...)
	fmt.Println("appended c:", c)

	i := make([]interface{}, 5)
	i[0] = 1
	fmt.Println(i)

	//j := []int{1,2,3}
	//fmt.Println(PrependItem(j, 4))

	//i = j  // !!! cannot use j (type []int) as type []interface {} in assignment
	fmt.Println(i)

	//k := []string{"aa", "bb", "cc"}
	//fmt.Println(PrependSlice(j, k))
}

func PrependItem(slice []interface{}, item interface{}) []interface{} {
	slice = append(slice, item)
	return slice
}

func PrependSlice(slice []interface{}, item []interface{}) []interface{} {
	for _, v := range item {
		slice = append(slice, v)
	}
	return slice
}

/*
func show(i int){
	fmt.Println(i)
}

func show(s string){
	fmt.Println(s)
}*/
