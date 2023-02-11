package account

import (
	// golang package
	"context"
	"database/sql"
	"time"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

//go:generate mockgen -source=./resource.go -destination=./resource_mock.go -package=account

// dbRepoProvider holds all methods from db repo that wil be used in account's resource.
type dbRepoProvider interface {
	// BeginTX will start a new transaction.
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)

	// Commit will commit the transaction.
	Commit(tx *sql.Tx) error

	// GetUserAccountByEmail will fetch user's information based of account's email.
	GetUserAccountByEmail(ctx context.Context, email string) (pgsql.Account, error)

	// InsertUserAccount will create a new entry in table user_account in database.
	InsertUserAccount(ctx context.Context, tx *sql.Tx, email, password string) error

	// Rollback will aborts the transaction.
	Rollback(tx *sql.Tx) error

	// UpdateUserAccount will update user's account information.
	UpdateUserAccount(ctx context.Context, tx *sql.Tx, param pgsql.UpdateUserAccountParam) error
}

// infraRepoProvider holds all methods from infra that will be needed in resource.
type infraRepoProvider interface {
	// GetConfig will get configuration that had been saved to memory.
	GetConfig() *configuration.AppConfig

	// JsonUnmarshal parses the JSON-encoded data and stores the result in the value pointed to by dest.
	JsonUnmarshal(input []byte, dest interface{}) error
}

// redisRepoProvider holds all methods from redis repo that wil be used in account's resource.
type redisRepoProvider interface {
	// Del will delete a key in redis.
	Del(ctx context.Context, key string) error

	// Get will get the value of a redis key.
	// First it will check whether the key exist or not.
	// If it exists, then it will return the value.
	// Otherwise, it returns empty string.
	Get(ctx context.Context, key string) (string, error)

	// Set will save the value of a key to redis.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

// AccountResourceParam holds all parameters needed to instantiate
// a new instance of Resource.
type AccountResourceParam struct {
	Cache redisRepoProvider
	Infra infraRepoProvider
	DB    dbRepoProvider
}

type Resource struct {
	cache redisRepoProvider
	infra infraRepoProvider
	db    dbRepoProvider
}

// NewResource will instantiate a new instance of Resource.
func NewResource(param AccountResourceParam) *Resource {
	return &Resource{
		cache: param.Cache,
		infra: param.Infra,
		db:    param.DB,
	}
}
