package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	"strconv"
	"crypto/rand"
	mathRand "math/rand"
	"time"
)

func main(){
	fmt.Println("This is main2.go from project2.")
	fmt.Println("-----------")

	//testFor()
	//testDefer()
	//testJSON()

	//testPrint()
	//testPointer()

	//testArgsPass()
	testArraySlice()

	testMathRand()
	testCryptoRand()
}


func testFor(){
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


func testDefer(){
	fmt.Println("Test defer start ...")
	defer fmt.Println("1st defer added.")

	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered:", 123)
		}
	}()

	for i := 0; i<3; i++ {
		fmt.Println("Simple loop: "+ strconv.Itoa(i))
	}

	j := 0
	for {
		defer fmt.Println("defered loop: " + strconv.Itoa(j))
		if j>5 {
			fmt.Println("Panic start ...")
			panic(fmt.Sprint("aaa"))
		}
		j++
	}
}

func testJSON(){
	var s []byte
	var _ error

	// string
	s, _ = json.Marshal("abcdefg")
	fmt.Println(string(s))

	// array
	arr := []uint{1,2,3,4,5,6,7}
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
		Age int
	}
	peter := Person{
		Name: "Peter",
		Age: 21,
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
		Page: 1,
		Fruits:"apple",
	}
	fmt.Println(res1D)
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
}

func testPrint(){
	fmt.Println("-------")

	v := "abcdef"

	fmt.Println(v)
	fmt.Printf("var1 value: %s, type: %T\n", v, v)

}

func testPointer(){
	fmt.Println("-------")

	s := "abc"
	sp := &s
	fmt.Printf("%v %v %T %v %T \n", s, sp ,sp, *sp, *sp)

	i := 123
	ip := &i
	fmt.Printf("%v %v %T \n", i, ip ,ip)

	// stuct
	type Person struct {
		Name string
		Age int
	}
	peter := Person{
		Name: "Peter",
		Age: 21,
	}
	r, _ := json.Marshal(peter)
	fmt.Printf("%v %T %s \n", peter, peter, r)

}

func testArgsPass(){
	a := 1
	b := true
	c := "3"
	printArgs(a, b, c, &a, &b, &c)
	fmt.Printf("%v %v %v\n", a, b, c)
}
func printArgs(a int, b bool, c string, ap *int, bp *bool, cp *string){
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
func testArraySlice(){
	originalArr := [6]int{1,2,3,4,5,6}
	fmt.Println(originalArr)

	arr := originalArr[:]
	fmt.Println("arr:", arr)
	//fmt.Println("arr == b", string(arr == originalArr)) // invalid operation: arr == originalArr (mismatched types []int and [6]int)

	s := arr[:2]
	fmt.Println("s:", s)

	arr[0] = 11
	fmt.Println("arr:",arr)

	printSlice(arr)
	fmt.Println("arr:",arr)
	fmt.Println("s:", s)


}
func printSlice(slice []int){
	slice[1] = 22
	fmt.Println("slice in sub func:", slice)
}

func testMathRand(){
	fmt.Println(getRand(100))
	fmt.Println(getRand(101))
}
func getRand(n int) (i int) {
	mathRand.Seed(time.Now().Unix())
	i = mathRand.Intn(n)
	return
}

func testCryptoRand(){
	c := 10
	b := make([]byte, c)
	i, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b, i)
}

