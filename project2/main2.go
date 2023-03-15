package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	mathRand "math/rand"
	"reflect"
	"strconv"
	"time"
)

func main() {
	fmt.Println("This is main2.go from project2.")
	fmt.Println("-----------")

	//testFor()
	//testDefer()
	//testJSON()

	//testPrint()
	// testPointer()

	//testArgsPass()
	testArraySlice()

	// testMathRand()
	// testCryptoRand()
}

func testFor() {
	// variable scope: block {}
	/*a := 1
	if true {
		a := 2
		fmt.Println(a)
	}
	fmt.Println(a)*/

	str := "123abc一二三"
	list := []rune(str)

	/*for i:=0; i<len(list); i++ {

		fmt.Println(string(list[i]))
	}*/

	for i, v := range list {
		fmt.Println(i, string(v))
	}
}

func testDefer() {
	fmt.Println("Test defer start ...")
	defer fmt.Println("1st defer added.")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", 123)
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("Simple loop: " + strconv.Itoa(i))
	}

	j := 0
	for {
		defer fmt.Println("defered loop: " + strconv.Itoa(j))
		if j > 5 {
			fmt.Println("Panic start ...")
			panic(fmt.Sprint("aaa"))
		}
		j++
	}
}

func testJSON() {
	var s []byte
	var _ error

	// string
	s, _ = json.Marshal("abcdefg")
	fmt.Println(string(s))

	// array
	arr := []uint{1, 2, 3, 4, 5, 6, 7}
	s, _ = json.Marshal(arr)
	fmt.Println(string(s))

	// map
	m := map[string]int{"Jason": 35, "Tom": 12, "Jerry": 12}
	s, _ = json.Marshal(m)
	fmt.Println(string(s), "len:", len(string(s)))

	fmt.Println("-------")

	// stuct
	type Person struct {
		Name string
		Age  int
	}
	peter := Person{
		Name: "Peter",
		Age:  21,
	}

	fmt.Println(reflect.TypeOf(peter))
	fmt.Println(peter)
	//fmt.Sprintf("%T", peter) // ???
	p, _ := json.Marshal(peter)
	fmt.Println(string(p))

	fmt.Println("-------")

	type Response1 struct {
		Page   int
		Fruits string
	}
	res1D := Response1{
		Page:   1,
		Fruits: "apple",
	}
	fmt.Println(res1D)
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
}

func testPrint() {
	fmt.Println("-------")

	v := "abcdef"

	fmt.Println(v)
	fmt.Printf("var1 value: %s, type: %T\n", v, v)

}

// stuct
type Person struct {
	Name string
	Age  int
}

func testPointer() {
	fmt.Println("-------")

	s := "abc"
	sp := &s
	fmt.Printf("%v %v %T %v %T \n", s, sp, sp, *sp, *sp)

	i := 123
	ip := &i
	fmt.Printf("%v %v %T \n", i, ip, ip)

	peter := Person{
		Name: "Peter",
		Age:  21,
	}
	r, _ := json.Marshal(peter)
	fmt.Printf("%v %T %s \n", peter, peter, r)

	fmt.Println("===")
	sp2 := printPointer1(sp)
	fmt.Printf("%p\n\n", sp2)

	fmt.Println("===")
	peter2 := printPointer2(&peter)
	fmt.Printf("%p\n\n", peter2)
}
func printPointer1(p *string) *string {
	fmt.Printf("%v %p\n", p, p)
	return p
}
func printPointer2(p *Person) *Person {
	fmt.Printf("%v %p\n", p, p)
	p.Age = 11
	fmt.Printf("%v %p\n", p, p)
	return p
}

func testArgsPass() {
	a := 1
	b := true
	c := "3"
	printArgs(a, b, c, &a, &b, &c)
	fmt.Printf("%v %v %v\n", a, b, c)
}
func printArgs(a int, b bool, c string, ap *int, bp *bool, cp *string) {
	fmt.Printf("%v %v %v\n", a, b, c)
	a = 2
	b = false
	c = "4"
	/*fmt.Printf("%v %v %v\n", a, b, c)
	*ap = 2
	*bp = false
	*cp = "4"*/
}

// slice itself is a reference, can be used as pointer
func testArraySlice() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s := arr[1:2]
	fmt.Printf("1, %p\n", s)
	// printSlice(&arr, &slice, &s)

	// arr2 := changeArray(arrp)
	// fmt.Printf("4, %p\n", arrp)
	// fmt.Printf("5, %p\n", &arr2)
	// printSlice(&arr, &slice, &s)

	// slice[0] = 11
	// printSlice(arr, slice, s)

	// s = append(s, 1, 2, 3)
	// printSlice(arr, slice, s)

	s = s[3:5]
	fmt.Printf("2, %p\n", s)

	// s = append(s, 8, 9)
	// printSlice(arr, slice, s)

	// persons := []Person{{Age: 1, Name: "a"}, {Age: 2, Name: "b"}, {Age: 3, Name: "c"}}
	// fmt.Printf("\n%p\n", persons)
	// persons2 := changePersons(persons)
	// fmt.Printf("%v %p\n", persons2, persons2)
}

func printSlice(arr [6]int, slice []int, s []int) {
	// fmt.Printf("%s %v %d %d %v\n", "s", s, len(*s), cap(*s), s)
	// fmt.Printf("%s %v %d %d %v\n", "slice", slice, len(*slice), cap(*slice), slice)
	fmt.Printf("%s %p %v\n\n", "arr", &arr, &arr)
}

func changeArray(arr *[6]int) *[6]int {
	fmt.Printf("2, %p\n", &arr)
	// fmt.Printf("%s %v %d %d %p\n", "changeArray, arr", *arr, len(arr), cap(arr), &arr)
	arr[5] = 123
	fmt.Printf("3, %p\n", &arr)
	// fmt.Printf("%s %v %d %d %p\n", "changeArray, arr", *arr, len(arr), cap(arr), &arr)
	return arr
}

func changePersons(s []Person) []Person {
	fmt.Printf("%v, %p\n", s, s)
	s[0].Age = 11
	fmt.Printf("%v, %p\n", s, s)
	s = append(s, Person{Age: 4, Name: "d"})
	fmt.Printf("%v, %p\n", s, s)
	return s
}

func testMathRand() {
	fmt.Println(getRand(100))
	fmt.Println(getRand(101))
}
func getRand(n int) (i int) {
	mathRand.Seed(time.Now().Unix())
	i = mathRand.Intn(n)
	return
}

func testCryptoRand() {
	c := 10
	b := make([]byte, c)
	i, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b, i)
}
