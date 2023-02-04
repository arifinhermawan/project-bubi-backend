package account

import (
	// golang package
	"context"
)

//go:generate mockgen -source=usecase.go -destination=usecase_mock.go -package=account

// accountServiceProvider holds all methods from account service that wil be used in account's usecase.
type accountServiceProvider interface {
	// CheckIsAccountExist will check whether an account is already exist by using email.
	CheckIsAccountExist(ctx context.Context, email string) (bool, error)

	// InsertUserAccount will create a new user account.
	InsertUserAccount(ctx context.Context, email, password string) (err error)
}

// AccountUsecaseParam holds all parameters needed to instantiate
// a new instance of Usecase.
type AccountUsecaseParam struct {
	Account accountServiceProvider
}

type UseCase struct {
	account accountServiceProvider
}

// NewUseCase will instantiate a new instance of UseCase.
func NewUseCase(param AccountUsecaseParam) *UseCase {
	return &UseCase{
		account: param.Account,
	}
}
