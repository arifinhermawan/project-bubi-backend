package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/usecase/account"
)

// UseCases holds all available usecases in bubi app.
type UseCases struct {
	account *account.UseCase
}

// NewUsecase will initialize a new instance of Usecases.
func NewUsecase(svc *Services) *UseCases {
	accountUseCaseParam := account.AccountUsecaseParam{
		Account: svc.account,
	}

	return &UseCases{
		account: account.NewUseCase(accountUseCaseParam),
	}
}
