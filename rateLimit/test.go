package main

import (
	"fmt"
	"math/rand"
	"rateLimit/util"
	"time"
)

func main() {
	t := time.NewTicker(time.Millisecond * 100)
	var limiter util.Limiter = testFixed()
	for now := range t.C {
		fmt.Println("---")
		rand.New(rand.NewSource(time.Now().UnixNano()))
		n := rand.Int63n(25)
		fmt.Println(now.Format("2006-01-02 15:04:05"), "request count:", n)
		fmt.Println(limiter.Get(uint64(n)))
	}
	defer t.Stop()
}

func testFixed() util.Limiter {
	return util.NewFixed(1000, 10)
}
