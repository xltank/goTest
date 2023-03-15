package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{1, 3, 5, 4, 2}
	//sort.Ints(intSlice)
	fmt.Println(intSlice)

	strSlice := []string{"aaa", "cba", "bca", "cab", "bac", "abc", "aab", "acb", "三二一", "二一三", "百千万", "诶比谁"}
	sort.Strings(strSlice)
	fmt.Println(strSlice)

	i := sort.Search(len(intSlice), func(i int) bool { return intSlice[i] == 5 })
	fmt.Println("in slice", intSlice, ", 5 is at index", i)
	fmt.Println(intSlice)

	sort.Reverse(sort.IntSlice(intSlice))
	fmt.Println(intSlice)

	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)
}
