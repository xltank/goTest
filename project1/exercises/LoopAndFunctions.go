package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can not Sqrt negtive number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	limit := 10

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else if x == 1 || x == 0 {
		return x, nil
	} else {
		a := 0.0
		b := x / 2
		fmt.Println(b)
		for i := 1; i < limit; i++ {
			a = b
			b = a + (x/a-a)/2
			fmt.Println(b)
		}
		return b, nil
	}
}

func main() {
	Sqrt(129374)
}
