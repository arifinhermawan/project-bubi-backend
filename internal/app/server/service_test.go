package server

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/service/account"
)

func TestNewService(t *testing.T) {
	mockRsc := &Resources{}
	mockInfra := &Infra{}

	want := &Services{
		account: account.NewService(account.AccountServiceParam{
			Infra: mockInfra,
			Rsc:   mockRsc.account,
		}),
	}

	got := NewService(mockRsc, mockInfra)
	assert.Equal(t, want, got)
}
