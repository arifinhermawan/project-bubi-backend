package account

import (
	// golang package
	"context"

	"github.com/arifinhermawan/bubi/internal/service/account"
)

//go:generate mockgen -source=usecase.go -destination=usecase_mock.go -package=account

// accountServiceProvider holds all methods from account service that wil be used in account's usecase.
type accountServiceProvider interface {
	// CheckPasswordCorrect will check whether user's password match with current password or not.
	CheckPasswordCorrect(ctx context.Context, email, password string) error

	// GenerateJWT will generate a new JWT for user if not exist.
	// If it already exist, it will get the value from cache.
	// If it's not exist, it will generate new token and save it to cache.
	GenerateJWT(ctx context.Context, userID int64, email string) (string, error)

	// GetUserAccountByEmail will check whether an account is already exist by using email.
	GetUserAccountByEmail(ctx context.Context, email string) (account.Account, error)

	// InsertUserAccount will create a new user account.
	InsertUserAccount(ctx context.Context, email, password string) (err error)

	// InvalidateJWT will delete user's cached JWT.
	InvalidateJWT(ctx context.Context, userID int64) error
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
