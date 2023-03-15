package main

import (
	//"golang.org/x/tour/reader"
	"strings"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(arr []byte) (int, error) {
	length := len(arr)
	str := strings.Repeat("A", length)
	return strings.NewReader(str).Read(arr)
}

/*
func main() {
    reader.Validate(MyReader{})
}
*/
