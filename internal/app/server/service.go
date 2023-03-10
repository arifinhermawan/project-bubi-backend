package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/service/account"
)

// Services holds all available services in bubi app.
type Services struct {
	account *account.Service
}

// NewService will initialize a new instance of Services.
func NewService(rsc *Resources, infra *Infra) *Services {
	accountServiceParam := account.AccountServiceParam{
		Rsc:   rsc.account,
		Infra: infra,
	}

	return &Services{
		account: account.NewService(accountServiceParam),
	}
}
