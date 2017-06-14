package main

import (
    "fmt"
    "encoding/json"
)

func stringify(obj interface{}) (r string, err error) {
    bytes, err := json.Marshal(obj)
    if err != nil {
        return
    }
    r = string(bytes)
    return
}

func parse(str string, obj interface{}) (err error){
    err = json.Unmarshal([]byte(str), obj)
    if err != nil {
        fmt.Println("Error when parsing string:", err)
    }
    return err
}

func main(){
    item1 := Item{123, "Jason", 1}
    item2 := Item{234, "Jack", 2}
    items := []Item{item1, item2}

    str, err := stringify(items)
    if err != nil {

    }
    fmt.Println("JSON:", str)


    //itemStr := `{"id":123,"name":"Jason"}`
    //itemsStr := `[{"id":123,"name":"Jason"},{"id":234,"name":"Jack"}]`
    escapedStr := "{\"id\":123,\"name\":\"Jason\"}"
    fmt.Println("String:", escapedStr)
    item := Item{}
    e := parse(escapedStr, &item)
    if e != nil {
        fmt.Println(err)
    }
    fmt.Println("Object:", item)


    intStr := `123456`
    var i int
    e2 := parse(intStr, &i)
    if e2 != nil {
        fmt.Println(err)
    }
    fmt.Println("int:", i)
}

type Item struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Role int `json:"-"`
}
