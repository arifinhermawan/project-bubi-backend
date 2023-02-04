package golang

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"
)

func TestNewGolang(t *testing.T) {
	want := &Golang{}
	got := NewGolang()
	assert.Equal(t, want, got)
}
