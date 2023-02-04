package server

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/server/account"
)

func TestNewHandler(t *testing.T) {
	got := NewHandler(&UseCases{}, &Infra{})

	usecases := &UseCases{}
	infra := &Infra{}

	accountHandlersParam := account.AccountHandlerParam{
		Account: usecases.account,
		Infra:   infra,
	}

	want := &Handlers{
		Account: account.NewHandler(accountHandlersParam),
	}

	assert.Equal(t, want, got)
}
