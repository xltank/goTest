/*
	Fixed window rate limiter.
*/

package util

import (
	"fmt"
	"time"
)

type Fixed struct {
	Limit  int // request limit
	Period int // window period, second
	count  int // request count
	start  time.Time
}

func NewFixed(limit int, period int) *Fixed {
	return &Fixed{
		Limit:  limit,
		Period: period,
		count:  0,
		start:  time.Now(),
	}
}

func (f *Fixed) Get(t time.Time) bool {
	now := time.Now()
	bound := f.start.Add(time.Duration(f.Period) * time.Second)
	if now.After(bound) {
		f.start = now
		f.count = 0
	}
	f.count += 1
	fmt.Println("Fixed:", f.count, f.Limit)
	return f.count < f.Limit
}
