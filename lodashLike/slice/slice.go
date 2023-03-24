package slice

import (
	"fmt"
	"reflect"
)

type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | string
}

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Add[T Addable](x ...T) T {
	var z T
	for i := 0; i < len(x); i++ {
		z += x[i]
	}
	return z
}

func Sum[T Num](x ...T) T {
	var z T
	for i := 0; i < len(x); i++ {
		z += x[i]
	}
	return z
}

// low efficient
func SumBy(s any, field string) (r float64) {
	fmt.Println("----- SumBy")
	re := reflect.ValueOf(s)
	if re.Type().Kind() != reflect.Slice {
		return 0
	}
	if re.Len() == 0 {
		return 0
	}
	// fmt.Println("re", re.Type())
	reFirst := re.Index(0)
	itemType := reFirst.Type()
	// fmt.Println("reFirst", itemType)
	if itemType.Kind() != reflect.Struct {
		return 0
	}
	for i := 0; i < re.Len(); i++ {
		fmt.Println("re.Index(i)", re.Index(i))
		reField := re.Index(i).FieldByName(field)
		fmt.Println("reItem", reField, reField.Type().Kind())
		value := float64(0)
		switch reField.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = float64(reField.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = float64(reField.Uint())
		case reflect.Float32, reflect.Float64:
			value = float64(reField.Float())
		}
		r += value
	}
	return r
}

func Map[T, F any](s []T, f func(t T) F) (r []F) {
	r = make([]F, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Compact2[T any](s []T) (r []T) {
	fmt.Println("----- Compact2")
	if len(s) == 0 {
		return s
	}
	r = make([]T, 0)
	// empty := new(T)
	/* for _, v := range s {
		if *empty != v {
			r = append(r, v)
		}
	} */

	return r
}

func Compact[T comparable](s []T) (r []T) {
	r = make([]T, 0)
	emp := new(T)
	for _, v := range s {
		if v == *emp {
			continue
		}
		r = append(r, v)
	}
	return r
}

// func Uniq[T comparable](s []T) (r []T) {

// }
