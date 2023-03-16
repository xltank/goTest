/*
	Sliding window rate limiter.
*/

package util

import (
	"time"
)

type Sliding struct {
	Limit  int // request limit
	Period int // window period, second
	list   []time.Time
	// sync.Mutex
}

func NewSliding(limit int, period int) *Sliding {
	return &Sliding{
		Limit:  limit,
		Period: period,
		list:   make([]time.Time, 0),
	}
}

func (s *Sliding) Get(t time.Time) bool {
	now := time.Now()
	start := now.Add(-time.Duration(s.Period) * time.Second)
	for _, v := range s.list {
		if v.Before(start) {
			s.list = s.list[1:]
		}
	}
	l := len(s.list)
	if l+1 > s.Limit {
		return false
	}
	s.list = append(s.list, t)
	return true
}
