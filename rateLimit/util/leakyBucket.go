package util

import (
	"fmt"
	"math"
	"time"
)

type LeakyBucket struct {
	Cap      int       // capacity of bucket
	Speed    int       // leaking(consuming) speed, per second
	count    int       // count of current req in bucket
	lastTime time.Time //
}

func (l *LeakyBucket) Get(t time.Time) bool {
	ellapsed := int(time.Since(l.lastTime).Seconds())
	if ellapsed >= 1 {
		l.count = int(math.Max(0, math.Min(float64(l.Cap), float64(l.count-l.Speed*int(ellapsed)))))
		l.lastTime = time.Now()
	}
	fmt.Printf("%v / %v:", l.count, l.Cap)
	if l.count >= l.Cap {
		return false
	}
	l.count++
	return true
}

func NewLeakyBucket(cap, speed int) *LeakyBucket {
	return &LeakyBucket{
		Cap:      cap,
		Speed:    speed,
		count:    0,
		lastTime: time.Now(),
	}
}
