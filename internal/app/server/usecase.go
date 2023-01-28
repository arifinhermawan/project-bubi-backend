package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/usecase/sample"
)

type UseCases struct {
	sample *sample.UseCase
}

func NewUsecase() *UseCases {
	return &UseCases{
		sample: sample.NewUsecase(),
	}
}
