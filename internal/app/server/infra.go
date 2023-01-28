package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

type configProvider interface {
	GetConfig() configuration.AppConfig
}

type Infra struct {
	Config configProvider
}

func NewInfra() *Infra {
	return &Infra{
		Config: configuration.NewConfiguration(),
	}
}
