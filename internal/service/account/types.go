package account

import (
	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
)

// Account is an entity representational of Account.
type Account entity.Account

// UpdateUserAccountParam represents parameters needed to update user's account
type UpdateUserAccountParam struct {
	FirstName    string
	LastName     string
	RecordPeriod int
	UserID       int64
}
