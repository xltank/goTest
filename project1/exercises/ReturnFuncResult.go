package main

import "fmt"

func main() {
	a, b := F1()
	fmt.Println(a, b)

	c, d := M1()
	fmt.Println(c, d)

	e, f := N1()
	fmt.Println(e, f)
}

func F2() (int, int) {
	return 1, 2
}

func F1() (int, int) {
	return F2()
}

var m = make(map[string]string)

func M1() (string, bool) {
	v, f := m["aa"]
	return v, f
}

// !!! not enough arguments to return, WHY ?
func N1() (string, bool) {
	return m["aa"]
}
