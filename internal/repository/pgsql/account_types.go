package pgsql

import (
	// golang package
	"database/sql"
)

// Account holds information about user's account
type Account struct {
	Email             string         `db:"email"`
	FirstName         sql.NullString `db:"first_name"`
	ID                int64          `db:"id"`
	LastName          sql.NullString `db:"last_name"`
	Password          string         `db:"password"`
	RecordPeriodStart int            `db:"record_period_start"`
}
