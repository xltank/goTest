package main

import "fmt"

type MyPoint struct {
	X int
	Y int
}

func printFuncValue(p MyPoint) {
	p.X = 1
	p.Y = 1
	fmt.Printf(" -> %v", p)
}

func printFuncPointer(pp *MyPoint) {
	pp.X = 1 // 实际上应该写做 (*pp).X，Golang 给了语法糖，减少了麻烦，但是也导致了 * 的不一致
	pp.Y = 1
	fmt.Printf(" -> %v", pp)
}

func (p MyPoint) printMethodValue() {
	p.X += 1
	p.Y += 1
	fmt.Printf(" -> %v", p)
}

func (pp *MyPoint) printMethodPointer() {
	pp.X += 1
	pp.Y += 1
	fmt.Printf(" -> %v", pp)
}

func main() {
	p := MyPoint{0, 0}
	pp := &MyPoint{0, 0}

	fmt.Printf("\n value to func(value): %v", p)
	printFuncValue(p)
	fmt.Printf(" --> %v", p)

	//printFuncValue(pp) // cannot use pp (type *MyPoint) as type MyPoint in argument to printFuncValue

	//printFuncPointer(p) // cannot use p (type MyPoint) as type *MyPoint in argument to printFuncPointer

	fmt.Printf("\n pointer to func(pointer): %v", pp)
	printFuncPointer(pp)
	fmt.Printf(" --> %v", pp)

	fmt.Printf("\n value to method(value): %v", p)
	p.printMethodValue()
	fmt.Printf(" --> %v", p)

	fmt.Printf("\n value to method(pointer): %v", p)
	p.printMethodPointer()
	fmt.Printf(" --> %v", p)

	fmt.Printf("\n pointer to method(value): %v", pp)
	pp.printMethodValue()
	fmt.Printf(" --> %v", pp)

	fmt.Printf("\n pointer to method(pointer): %v", pp)
	pp.printMethodPointer()
	fmt.Printf(" --> %v", pp)
}
