package main

import (
	"fmt"
	"math/rand"
	"rateLimit/util"
	"time"
)

func main() {
	// test1(util.NewFixed(1000, 10))
	// test2(util.NewSliding(1000, 10))
	// test2(util.NewLeakyBucket(1000, 100))
	test2(util.NewTokenBucket(1000, 100))
}

/* func test1(limiter util.Limiter) {
	t := time.NewTicker(time.Millisecond * 100)
	for now := range t.C {
		fmt.Println("---")
		rand.New(rand.NewSource(time.Now().UnixNano()))
		n := rand.Int63n(25)
		fmt.Println(now.Format("2006-01-02 15:04:05"), "request count:", n)
		fmt.Println(limiter.Get(int(n)))
	}
	defer t.Stop()
} */

func test2(limiter util.Limiter) {
	c := make(chan time.Time)
	for i := 0; i < 20; i++ {
		go (func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			ticker := time.NewTicker((time.Millisecond * 50))
			for t := range ticker.C {
				rn := r.Int31n(10)
				if rn < 5 {
					// fmt.Printf("go [%v], rand: %v\n", i, rn)
					c <- t
				}
			}
		})()
	}
	for t := range c {
		// fmt.Println(t.Format("2006-01-02 15:04:05.999999999"))
		fmt.Println(limiter.Get(t))
	}
}
