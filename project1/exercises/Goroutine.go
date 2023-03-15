package main

import (
	"fmt"
	"time"
)

func main() {

	/*for i:=0; i<20; i++ {
	    go func(){
	        fmt.Println(i)
	    }()
	}*/

	/*arr := []int{1,2,3,4,5,6,7,8,9,10}
	  for v := range arr {
	      go func(){
	          fmt.Println(v)
	      }()
	  }*/

	c := make(chan int)
	go func() { fmt.Println(<-c) }()
	c <- 1

	<-time.After(time.Second)
}
