package common

import "fmt"

func Print(s string) string {
	fmt.Println(s)
	utilsVar = 2
	privateFunc()
	return s
}

var (
	privateVar int = 1
)
