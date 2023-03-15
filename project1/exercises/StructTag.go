package main

import (
	"fmt"
	"reflect"
)

func main() {
	item := Item{123, "Jason", 1}
	t := reflect.TypeOf(item)
	f, _ := t.FieldByName("Role")
	fmt.Println(t.Name(), t.Kind(), t.NumField(), t.Field(0), f, f.Tag, t.Size(), t.Field(0).Tag.Get("json"))
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role int    `json:"-"`
}
