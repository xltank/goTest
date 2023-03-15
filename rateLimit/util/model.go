package util

type Limiter interface {
	Get(n uint64) bool
}
