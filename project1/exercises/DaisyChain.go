package main

import "fmt"

func pass(left, right chan int) {
	left <- 1 + <-right
}

/*func main(){
    const n = 50
    leftmost := make(chan int)
    right := leftmost
    left := leftmost

    for i := 0; i< n; i++ {
        right = make(chan int)
        // the chain is constructed from the end
        go pass(left, right) // the first goroutine holds (leftmost, new chan)
        left = right         // the second and following goroutines holds (last right chan, new chan)
    }
    go func(c chan int){ c <- 1}(right)
    fmt.Println("sum:", <- leftmost)
}*/

// this chain is constructed from the start
func main() {
	const n = 50

	start := make(chan int)
	right := start

	for i := 0; i < n-1; i++ {
		left := make(chan int)
		go pass(left, right)
		right = left
	}
	result := make(chan int)
	go pass(result, right)

	go func(c chan int) { c <- 1 }(start)
	fmt.Println(<-result)
}
