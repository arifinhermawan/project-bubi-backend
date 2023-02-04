package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/server/account"
)

// Handlers holds all available handlers in bubi app.
type Handlers struct {
	Account *account.Handler
}

// NewHandler initialize new instance of Handlers.
func NewHandler(usecases *UseCases, infra *Infra) *Handlers {
	accountHandlerParam := account.AccountHandlerParam{
		Account: usecases.account,
		Infra:   infra,
	}

	return &Handlers{
		Account: account.NewHandler(accountHandlerParam),
	}
}
