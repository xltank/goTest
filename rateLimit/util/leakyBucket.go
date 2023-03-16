package util

import (
	"fmt"
	"math"
	"time"
)

type Leaky struct {
	Cap        int       // capacity of bucket
	Speed      int       // leaking(consuming) speed, per second
	tokenCount int       // count of current token in bucket
	lastTime   time.Time //
}

func (l *Leaky) Get(t time.Time) bool {
	ellapsed := int(time.Since(l.lastTime).Seconds())
	if ellapsed >= 1 {
		l.tokenCount = int(math.Min(float64(l.Cap), float64(l.tokenCount+l.Speed*int(ellapsed))))
		l.lastTime = time.Now()
	}
	fmt.Printf("%v / %v:", l.tokenCount, l.Cap)
	if l.tokenCount == 0 {
		return false
	}
	l.tokenCount--
	return true
}

func NewLeaky(cap, speed int) *Leaky {
	return &Leaky{
		Cap:        cap,
		Speed:      speed,
		tokenCount: cap,
		lastTime:   time.Now(),
	}
}
