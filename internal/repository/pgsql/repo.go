package pgsql

import (
	// external package
	"context"
	"database/sql"
	"time"

	// external package
	"github.com/jmoiron/sqlx"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

var (
	funcSQLXNamed = sqlx.Named
)

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=pgsql

// infraProvider holds all methods from infra that is needed in repository.
type infraProvider interface {
	// GetConfig will get configuration that had been saved to memory.
	GetConfig() *configuration.AppConfig

	// GetTimeGMT7 will get current time in GMT+7
	GetTimeGMT7() time.Time
}

// psqlProvider holds all methods from sql that is needed in repository.
type psqlProvider interface {
	// BeginTx starts a transaction.
	//
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the sql package will roll back
	// the transaction. Tx.Commit will return an error if the context provided to
	// BeginTx is canceled.
	//
	// The provided TxOptions is optional and may be nil if defaults should be used.
	// If a non-default isolation level is used that the driver doesn't support,
	// an error will be returned.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// GetContext using this Conn.
	// Any placeholder parameters are replaced with supplied args.
	// An error is returned if the result set is empty.
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Rebind a query within a Conn's bindvar type.
	Rebind(query string) string
}

// DBRepoParam holds all parameters needed to instansiate new
// DBRepository instance
type DBRepoParam struct {
	DB    psqlProvider
	Infra infraProvider
}

type DBRepository struct {
	db    psqlProvider
	infra infraProvider
}

// NewDBRepository will create a new instance of DBRepository.
func NewDBRepository(param DBRepoParam) *DBRepository {
	return &DBRepository{
		db:    param.DB,
		infra: param.Infra,
	}
}

// BeginTX will start a new transaction.
func (repo *DBRepository) BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return repo.db.BeginTx(ctx, options)
}

// Commit will commit the transaction.
func (repo *DBRepository) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

// Rollback will aborts the transaction.
func (repo *DBRepository) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
