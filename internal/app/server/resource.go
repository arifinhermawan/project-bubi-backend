package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/service/sample"
)

type Resources struct {
	sample *sample.Resource
}

func NewResource() *Resources {
	return &Resources{
		sample: sample.NewResource(),
	}
}
