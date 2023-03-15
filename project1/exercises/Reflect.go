package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	prod := Producer{001, "ABC Press", "BeiJing"}
	item := Item2{111, "Book", 12.34, time.Date(2017, 6, 8, 18, 03, 12, 345, time.UTC), prod}
	t := reflect.TypeOf(item)
	fmt.Println("Type of item is:", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i))
	}

	itemDefault := reflect.New(t)
	fmt.Println(itemDefault)
	fmt.Println(reflect.Zero(t))

	f := func(a int, b string, c bool) {
		fmt.Println(a, b, c)
	}
	fmt.Println(reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("aaa"), reflect.ValueOf(false)}))

	fmt.Println("-------------")

	blank := Item2{}
	ScanStruct(&blank)
}

type Item2 struct {
	Id           int
	Name         string
	Price        float32
	ProducedDate time.Time
	Producer     Producer
}

type Producer struct {
	Id      int
	Name    string
	Address string
}

func Scan(args ...interface{}) {
	fmt.Println(args...)
}

func ScanStruct(sp interface{}) error {
	vp := reflect.ValueOf(sp)
	if vp.Kind() != reflect.Ptr {
		return errors.New("Must be a Pointer.")
	}
	v := vp.Elem()

	fieldPointers := []interface{}{}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fp := f.Addr().Interface()
		fieldPointers = append(fieldPointers, fp)
	}
	fmt.Println(fieldPointers, len(fieldPointers))

	Scan(fieldPointers...)

	return nil
}
