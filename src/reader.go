package main

import (
	"fmt"
	"io"
)

// StringReader is a type created do implement a io.reader example
type StringReader struct{}

func (r StringReader) Read(p []byte) (int, error) {
	p[0] = 'A'
	p[1] = 'B'
	p[2] = 'C'
	p[3] = 'D'

	return len(p), nil
}

func readString(r io.Reader) string {
	p := make([]byte, 4)
	r.Read(p)
	return string(p)
}

func main() {
	reader := StringReader{}
	fmt.Println(readString(reader))
}
