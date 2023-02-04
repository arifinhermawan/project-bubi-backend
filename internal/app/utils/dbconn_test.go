package utils

import (
	// golang package
	"testing"

	// external package
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

func TestInitDBConn(t *testing.T) {
	sqlxOpenOri := sqlxOpen
	defer func() {
		sqlxOpen = sqlxOpenOri
	}()

	type args struct {
		cfg configuration.DatabaseConfig
	}
	tests := []struct {
		name        string
		args        args
		mockSqlOpen func(driverName, dataSourceName string) (*sqlx.DB, error)
		want        *sqlx.DB
		wantErr     error
	}{
		{
			name: "when_sqlOpen_error_then_return_error",
			args: args{},
			mockSqlOpen: func(driverName, dataSourceName string) (*sqlx.DB, error) {
				return nil, assert.AnError
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_db_and_nil",
			args: args{},
			mockSqlOpen: func(driverName, dataSourceName string) (*sqlx.DB, error) {
				return &sqlx.DB{}, nil
			},
			want: &sqlx.DB{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sqlxOpen = test.mockSqlOpen
			got, err := InitDBConn(test.args.cfg)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
