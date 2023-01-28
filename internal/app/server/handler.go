package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/server/sample"
)

type Handlers struct {
	Sample *sample.Handler
}

func NewHandler() *Handlers {
	return &Handlers{
		Sample: sample.NewHandler(),
	}
}
