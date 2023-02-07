package server

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
	"github.com/arifinhermawan/bubi/internal/repository/redis"
	"github.com/arifinhermawan/bubi/internal/service/account"
)

func TestNewResource(t *testing.T) {
	mockDB := &pgsql.DBRepository{}
	mockCache := &redis.RedisRepository{}
	mockInfra := &Infra{}

	want := &Resources{
		account: account.NewResource(account.AccountResourceParam{
			DB:    mockDB,
			Cache: mockCache,
			Infra: mockInfra,
		}),
	}

	got := NewResource(ResourceParam{
		DB:    mockDB,
		Cache: mockCache,
		Infra: mockInfra,
	})

	assert.Equal(t, want, got)
}
