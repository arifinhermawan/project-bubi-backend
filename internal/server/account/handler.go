package account

import (
	// golang package
	"context"
	"io"

	// internal package
	"github.com/arifinhermawan/bubi/internal/usecase/account"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=account

// accountUCManager holds all methods served by usecase account that will be needed by account handler.
type accountUCManager interface {
	// LogIn handles the log in process for a user.
	// It will check the existence of a user first.
	// If it exist, then it will continue the log in process.
	LogIn(ctx context.Context, email, password string) (string, error)

	// LogOut handles the log out process for a user.
	LogOut(ctx context.Context, userID int64) error

	// UpdateUserAccount will update information of a user account.
	// Field that will be updated are: first_name, last_name, and record_period.
	UpdateUserAccount(ctx context.Context, param account.UpdateUserAccountParam) error

	// UpdatePassword will update user's password.
	// It will check whether the old password correct or not.
	// If it correct, then it will continue the update password process.
	UpdatePassword(ctx context.Context, param account.UpdatePasswordParam) error

	// UserSignUp will process the creation of user account.
	// Before creating a new account, it'll check whether that account exist or not.
	// If it's a new account, then it'll create a new user account.
	UserSignUp(ctx context.Context, email, password string) error
}

// infraProvider holds all methods served by infra that will be needed by account handler.
type infraProvider interface {
	// JsonUnmarshal parses the JSON-encoded data and stores the result in the value pointed to by dest.
	JsonUnmarshal(input []byte, dest interface{}) error

	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF. Because ReadAll is
	// defined to read from src until EOF, it does not treat an EOF from Read
	// as an error to be reported.
	ReadAll(input io.Reader) ([]byte, error)
}

// AccountHandlerParam holds all parameters needed to instantiate a new account Handler.
type AccountHandlerParam struct {
	Account accountUCManager
	Infra   infraProvider
}

type Handler struct {
	account accountUCManager
	infra   infraProvider
}

// NewHandler instantiate a new instance of Handler.
func NewHandler(param AccountHandlerParam) *Handler {
	return &Handler{
		account: param.Account,
		infra:   param.Infra,
	}
}
