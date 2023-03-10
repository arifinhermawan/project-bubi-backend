package server

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
	"github.com/arifinhermawan/bubi/internal/repository/redis"
	"github.com/arifinhermawan/bubi/internal/service/account"
)

// Resources holds all available resources in bubi app.
type Resources struct {
	account *account.Resource
}

// ResourceParam represents parameters needed to initialize Resources.
type ResourceParam struct {
	DB    *pgsql.DBRepository
	Cache *redis.RedisRepository
	Infra *Infra
}

// NewResource will initialize a new instance of Resources.
func NewResource(param ResourceParam) *Resources {
	accountResourceParam := account.AccountResourceParam{
		Cache: param.Cache,
		Infra: param.Infra,
		DB:    param.DB,
	}

	return &Resources{
		account: account.NewResource(accountResourceParam),
	}
}
