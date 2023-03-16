package util

import "time"

type Limiter interface {
	Get(t time.Time) bool
}
