package main

import (
	"fmt"
)

func main() {
	var h Human

	s := Student{Grade: 1, Major: "English", Human: Human{Name: "Jason", Age: 12, Being: Being{IsLive: true}}}
	fmt.Println("student:", s)
	fmt.Println("student:", s.Name, ", isLive:", s.IsLive, ", age:", s.Age, ", grade:", s.Grade, ", major:", s.Major)

	//h = s // cannot use s (type Student) as type Human in assignment
	fmt.Println(h)

	//Heal(s) // cannot use s (type Student) as type Being in argument to Heal
	Heal(s.Human.Being) // true

	s.Drink()
	s.Eat()

	v1, b := interface{}(s).(Car)
	fmt.Println(v1, b)

	v2, b := interface{}(s).(Being)
	fmt.Println(v2, b)

	v3, b := interface{}(s).(Human)
	fmt.Println(v3, b)

	v4, b := interface{}(s).(Student)
	fmt.Println(v4, b)

	v5, b := interface{}(s).(IDrink)
	fmt.Println(v5, b)

	v6, b := interface{}(s).(IEat)
	fmt.Println(v6, b)

	v7, b := interface{}(s).(IMove)
	fmt.Println(v7, b)

	v8, b := interface{}(s).(int)
	fmt.Println(v8, b)

	s1 := Student{Grade: 1, Major: "English", Human: Human{Name: "Jason", Age: 12, Being: Being{IsLive: true}}}
	s2 := Student{Grade: 1, Major: "English", Human: Human{Name: "Tom", Age: 13, Being: Being{IsLive: true}}}
	s3 := Student{Grade: 1, Major: "English", Human: Human{Name: "Mike", Age: 14, Being: Being{IsLive: true}}}
	t1 := Teacher{Grade: 1, Major: "English", Salary: 2000, Human: Human{Name: "Michael", Age: 34, Being: Being{IsLive: true}}}
	t2 := Teacher{Grade: 1, Major: "English", Salary: 3000, Human: Human{Name: "Tony", Age: 31, Being: Being{IsLive: true}}}
	t3 := Teacher{Grade: 1, Major: "English", Salary: 4000, Human: Human{Name: "Ivy", Age: 40, Being: Being{IsLive: true}}}
	drinkers := []IDrink{s1, s2, s3, t1, t2, t3}

	for _, v := range drinkers {
		switch t := v.(type) {
		case Student:
			fmt.Println(t.Name, "is a Student, he/she needs more homework.")
		case Teacher:
			fmt.Println(t.Name, "is a Teacher, he/she needs more jobs.")
		default:
			fmt.Println("Invalid Human being:", t)
		}

	}
}

type Car struct {
	Color     string
	SeatCount int
}

type Being struct {
	IsLive bool
}

type Human struct {
	Being
	Name string
	Age  int
}

func (h Human) Eat() {
	fmt.Println("human eating...")
	h.Drink()
}

func (h Human) Drink() {
	fmt.Println("human drinking...")
}

type Student struct {
	Human
	School string
	Grade  int
	Major  string
}

func (s Student) Drink() {
	fmt.Println("student drinking...")
}

type Teacher struct {
	Human
	School string
	Grade  int
	Major  string
	Salary int
}

func (s Teacher) Drink() {
	fmt.Println("teacher drinking...")
}

type IEat interface {
	Eat()
}

type IMove interface {
	Move()
}

type IDrink interface {
	Drink()
}

func Heal(b Being) {
	fmt.Println(b.IsLive)
}
