package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/service/sample"
)

type Services struct {
	sample *sample.Service
}

func NewService() *Services {
	return &Services{
		sample: sample.NewService(),
	}
}
