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

func SumBy[T any](s []T, field string) (r float64) {
	fmt.Println("----- SumBy")
	if len(s) == 0 {
		return 0
	}
	itemType := reflect.ValueOf(s[0]).Type()
	if itemType.Kind() != reflect.Struct {
		return 0
	}
	for _, v := range s {
		reField := reflect.ValueOf(v).FieldByName(field)
		if !reField.IsValid() {
			return 0
		}
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

func FindFirstIndex[T comparable](s []T, item T) (index int) {
	index = -1
	for i, v := range s {
		if item == v {
			return i
		}
	}
	return
}

func Uniq[T comparable](s []T) (r []T) {
	r = make([]T, 0)
	for _, v := range s {
		index := FindFirstIndex(r, v)
		if index == -1 {
			r = append(r, v)
		}
	}

	return
}
