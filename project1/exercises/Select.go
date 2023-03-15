package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("from goroutine ...", time.Now())
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		fmt.Println("quit ready")
		quit <- 0
	}()
	fmt.Println("start fibonacci ...", time.Now())
	fibonacci(c, quit)
}
