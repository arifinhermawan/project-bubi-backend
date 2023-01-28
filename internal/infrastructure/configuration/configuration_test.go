package configuration

import (
	// golang package
	"sync"
	"testing"

	// external package
	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	want := &Configuration{
		Config:           AppConfig{},
		doLoadConfigOnce: &sync.Once{},
	}

	got := NewConfiguration()
	assert.Equal(t, want, got)
}
