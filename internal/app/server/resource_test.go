package server

import (
	// golang package
	"testing"

	// external package
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
	"github.com/arifinhermawan/bubi/internal/service/account"
)

func TestNewResource(t *testing.T) {
	mockDB := &pgsql.DBRepository{}

	want := &Resources{
		account: account.NewResource(account.AccountResourceParam{
			DB: mockDB,
		}),
	}

	got := NewResource(ResourceParam{
		DB: mockDB,
	})

	assert.Equal(t, want, got)
}
