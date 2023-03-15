package main

import (
	"fmt"
	"strconv"
	"time"
)

func push(c chan int) {
	for i := 1; i < 11; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(">>>", i, "from goroutine", getNow())
		c <- i
	}
	close(c) // return 0, false from channel
}

func main() {
	c := make(chan int, 2)
	go push(c)

	for {
		//time.Sleep(1800 * time.Millisecond)
		x, b := <-c
		fmt.Println("< got result from goroutine", x, b, time.Now())
		if !b {
			break
		}
	}
}

type MyTime struct {
	time.Time
}

/*func New() MyTime {
    t = time.Now()
    return t
}*/

func getNow() string {
	t := time.Now()
	return strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())
}
