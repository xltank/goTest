package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := &MyStruct2{v: 0}

	go f(s)
	go f(s)
	go f(s)
	go f(s)

	time.Sleep(time.Second)
	fmt.Println(" -> ", s.v)
}

type MyStruct2 struct {
	v   int
	mux sync.Mutex
}

func (s *MyStruct2) Inc() int {
	s.mux.Lock()
	s.v++
	fmt.Println(s.v)
	s.mux.Unlock()
	return s.v
}

func f(s *MyStruct2) {
	for i := 0; i < 30; i++ {
		s.Inc()
	}
}
