package server

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/usecase/account"
)

func TestNewUsecase(t *testing.T) {
	mockSvc := &Services{}

	want := &UseCases{
		account: account.NewUseCase(account.AccountUsecaseParam{
			Account: mockSvc.account,
		}),
	}

	got := NewUsecase(mockSvc)
	assert.Equal(t, want, got)
}
