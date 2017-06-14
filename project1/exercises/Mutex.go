package main

import (
    "sync"
    "fmt"
    "time"
)

func main (){
    s := &MyStruct{v:0}

    go f(s)
    go f(s)
    go f(s)
    go f(s)

    time.Sleep(time.Second)
    fmt.Println(" -> ",s.v)
}


type MyStruct struct {
    v int
    mux sync.Mutex
}


func (s *MyStruct) Inc() int{
    s.mux.Lock()
    s.v ++
    fmt.Println(s.v)
    s.mux.Unlock()
    return s.v
}


func f(s *MyStruct){
    for i:=0; i<30; i++ {
        s.Inc()
    }
}