package reader

import (
	// golang package
	"io"
)

type Reader struct{}

// NewReader initialize new instance of reader
func NewReader() *Reader {
	return &Reader{}
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (r *Reader) ReadAll(input io.Reader) ([]byte, error) {
	return io.ReadAll(input)
}
