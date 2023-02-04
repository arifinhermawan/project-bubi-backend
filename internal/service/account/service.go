package account

import (
	// golang package
	"context"

	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
)

//go:generate mockgen -source=./service.go -destination=./service_mock.go -package=account

// resourceProvider holds all methods from resource that wil be used in account's service.
type resourceProvider interface {
	// GetUserAccountByEmailFromDB will fetch user's information based of account's email.
	GetUserAccountByEmailFromDB(ctx context.Context, email string) (entity.Account, error)

	// InsertUserAccountToDB will create a new entry of user account in database.
	InsertUserAccountToDB(ctx context.Context, email, password string) error
}

// AccountServiceParam holds all parameters needed to instantiate
// a new instance of Service.
type AccountServiceParam struct {
	Rsc resourceProvider
}

type Service struct {
	rsc resourceProvider
}

// NewService will instantiate a new instance of Service.
func NewService(param AccountServiceParam) *Service {
	return &Service{
		rsc: param.Rsc,
	}
}
