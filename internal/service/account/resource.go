package account

import (
	// golang package
	"context"
	"database/sql"

	// internal package
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

//go:generate mockgen -source=./resource.go -destination=./resource_mock.go -package=account

// dbRepoProvider holds all methods from repo that wil be used in account's resource.
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
}

// AccountResourceParam holds all parameters needed to instantiate
// a new instance of Resource.
type AccountResourceParam struct {
	DB dbRepoProvider
}

type Resource struct {
	db dbRepoProvider
}

// NewResource will instantiate a new instance of Resource.
func NewResource(param AccountResourceParam) *Resource {
	return &Resource{
		db: param.DB,
	}
}
