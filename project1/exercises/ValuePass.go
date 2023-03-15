package main

import "fmt"

func main() {

	/*x := 1
	  y := 2;
	  xp := &x;
	  y = *xp;
	  x = xp; // unlike those in C: cannot use xp (type *int) as type int in assignment
	  *xp = 3;
	  fmt.Println(x, y, xp, *xp)
	  return*/

	fmt.Println("\n--- int")

	i := 0
	fmt.Printf("%v", i)
	passInt(i)
	fmt.Printf(" --> %v", i)

	fmt.Println("\n--- int pointer")

	ip := &i
	fmt.Printf("%v", ip)
	passIntPointer(ip)
	fmt.Printf(" --> %v", ip)

	fmt.Println("\n--- complex")

	fmt.Println("\n--- string")

	str := "0"
	fmt.Printf("%v", str)
	passString(str)
	fmt.Printf(" --> %v", str)

	fmt.Println("\n--- string pointer")

	strP := &str
	fmt.Printf("%v", strP)
	passStringPointer(strP)
	fmt.Printf(" --> %v", strP)

	fmt.Println("\n--- bool")

	b := false
	fmt.Printf("%v", b)
	passBool(b)
	fmt.Printf(" --> %v", b)

	fmt.Println("\n--- bool pointer")

	bp := &b
	fmt.Printf("%v", bp)
	passBoolPointer(bp)
	fmt.Printf(" --> %v", bp)

	fmt.Println("\n--- struct")

	s := Class{0, "a"}
	fmt.Printf("%v", s)
	s2 := passStruct(s)
	fmt.Printf(" --> %v", s)
	fmt.Println("\n s==s2:", s == s2)

	fmt.Println("\n--- struct pointer")

	sp := &Class{0, "a"}
	fmt.Printf("%v", sp)
	passStructPointer(sp)
	fmt.Printf(" --> %v", sp)

	fmt.Println("\n--- slice")

	/**
	  A slice does not store any data, it just describes a section of an underlying array.
	  Changing the elements of a slice modifies the corresponding elements of its underlying array.
	  Other slices that share the same underlying array will see those changes.
	*/
	sl := []int{0, 1, 2, 3, 4}
	fmt.Printf("%v", sl)
	passSlice(sl)
	fmt.Printf(" --> %v", sl)

	fmt.Println("\n--- slice pointer")

	slp := &[]int{0, 1, 2, 3, 4}
	fmt.Printf("%v", slp)
	passSlicePointer(slp)
	fmt.Printf(" --> %v", slp)

	fmt.Println("\n--- map")

	mm := map[string]string{
		"first":  "First one.",
		"second": "Second one.",
	}
	fmt.Printf("%v", mm)
	passMap(mm)
	fmt.Printf(" --> %v", mm)

	fmt.Println("\n--- map pointer")

	mmp := &map[string]string{
		"first":  "First one.",
		"second": "Second one.",
	}
	fmt.Printf("%v", mmp)
	passMapPointer(mmp)
	fmt.Printf(" --> %v", mmp)
}

func passInt(i int) {
	i = 1
	fmt.Printf(" -> %v", i)
}

func passIntPointer(ip *int) {
	*ip = 1
	fmt.Printf(" -> %v", *ip)
}

func passString(str string) {
	str = "1"
	fmt.Printf(" -> %v", str)
}

func passStringPointer(strP *string) {
	//strP = "1";
	fmt.Printf(" -> %v", strP)
}

func passBool(bp bool) {
	bp = true
	fmt.Printf(" -> %v", bp)
}

func passBoolPointer(b *bool) {
	//b = true;
	fmt.Printf(" -> %v", b)
}

type Class struct {
	Index int
	Value string
}

/*
func (c Class) output(){
    fmt.Println("Class c:", c)
}
*/

func passStruct(s Class) Class {
	s.Index = 1
	s.Value = "b"
	fmt.Printf(" -> %v", s)
	return s
}

func passStructPointer(sp *Class) {
	sp.Index = 1
	sp.Value = "b"
	fmt.Printf(" -> %v", sp)
}

func passSlice(s []int) {
	s[0] = 123
	fmt.Printf(" -> %v", s)
}

func passSlicePointer(sp *[]int) {
	//sp[0] = 123
	fmt.Printf(" -> %v", sp)
}

func passMap(mm map[string]string) {
	mm["latest"] = "The latest one."
	fmt.Printf(" -> %v", mm)
}

func passMapPointer(mmp *map[string]string) {
	// !!! invalid operation: mmp["latest"] (type *map[string]string does not support indexing)
	//mmp["latest"] = "The latest one."
	fmt.Printf(" -> %v", mmp)
}
