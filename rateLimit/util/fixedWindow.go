/*
	Fixed window rate limiter.
*/

package util

import (
	"fmt"
	"time"
)

type Fixed struct {
	Limit  uint64 // request limit
	Period uint64 // window period, second
	Count  uint64 // request count
	start  time.Time
}

func NewFixed(limit uint64, period uint64) Fixed {
	return Fixed{
		Limit:  limit,
		Period: period,
		Count:  0,
		start:  time.Now(),
	}
}

func (f *Fixed) Get(n uint64) bool {
	if n == 0 {
		n = 1
	}
	now := time.Now()
	bound := f.start.Add(time.Duration(f.Period) * time.Second)
	if now.After(bound) {
		f.start = now
		f.Count = 0
	}
	f.Count += n
	fmt.Println("Fixed:", f.Count, f.Limit)
	return f.Count < f.Limit
}
