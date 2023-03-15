package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	//go func (){
	ch <- 1
	//}()

	go func() {
		fmt.Println(<-ch)
	}()

	/*for i:=0; i<100; i++{
	      ch <- i
	  }

	  go show(ch)
	  go show(ch)
	  go show(ch)
	  fmt.Println("It's my turn", <- ch)*/

	<-time.After(time.Second * 1)
	fmt.Println("The end.")
}

func show(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
