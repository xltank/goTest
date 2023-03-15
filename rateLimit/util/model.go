package util

type Limiter interface {
	Get(n int) bool
}
