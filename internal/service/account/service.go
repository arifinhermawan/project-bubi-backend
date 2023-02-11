package account

import (
	// golang package
	"context"
	"time"

	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

//go:generate mockgen -source=./service.go -destination=./service_mock.go -package=account

// resourceProvider holds all methods from resource that wil be used in account's service.
type resourceProvider interface {
	// DeleteJWTInCache will delete JWT of a user.
	DeleteJWTInCache(ctx context.Context, userID int64) error

	// GetJWTFromCache will fetch JWT of a user from cache.
	GetJWTFromCache(ctx context.Context, userID int64) (string, error)

	// GetUserAccountByEmailFromDB will fetch user's information based of account's email.
	GetUserAccountByEmailFromDB(ctx context.Context, email string) (entity.Account, error)

	// InsertUserAccountToDB will create a new entry of user account in database.
	InsertUserAccountToDB(ctx context.Context, email, password string) error

	// SetJWTToCache will save jwt of a user in cache
	SetJWTToCache(ctx context.Context, userID int64, jwt string) error

	// UpdateUserAccountInDB will update user's account based on the given parameter.
	UpdateUserAccountInDB(ctx context.Context, param UpdateUserAccountParam) error
}

// infraProvider holds all methods from infra that will be needed in resource.
type infraProvider interface {
	// GetConfig will get configuration that had been saved to memory.
	GetConfig() *configuration.AppConfig

	// GetTimeGMT7 will get current time in GMT+7
	GetTimeGMT7() time.Time
}

// AccountServiceParam holds all parameters needed to instantiate
// a new instance of Service.
type AccountServiceParam struct {
	Infra infraProvider
	Rsc   resourceProvider
}

type Service struct {
	infra infraProvider
	rsc   resourceProvider
}

// NewService will instantiate a new instance of Service.
func NewService(param AccountServiceParam) *Service {
	return &Service{
		infra: param.Infra,
		rsc:   param.Rsc,
	}
}
