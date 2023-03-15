package main

import (
	"fmt"
	"runtime"
	"strings"
)

func callers() []uintptr {
	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	st := pcs[0:n]
	return st
}

type trace struct {
	m string
	s []uintptr
}

func (e *trace) Error() string {
	var b strings.Builder
	b.WriteString(e.m)
	b.WriteString("\n\n")
	b.WriteString("Traceback:")
	for _, pc := range e.s {
		fn := runtime.FuncForPC(pc)
		b.WriteString("\n")
		f, n := fn.FileLine(pc)
		b.WriteString(fmt.Sprintf("%s:%d", f, n))
	}
	return b.String()
}

// NewTrace creates a simple traceable error.
func NewTrace(message string) error {
	return &trace{m: message, s: callers()}
}

func f() error {
	return NewTrace("ooops")
}

func main() {
	fmt.Println(f())
}
