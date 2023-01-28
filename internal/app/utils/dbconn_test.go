package utils

import (
	// golang package
	"database/sql"
	"testing"

	// external package
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

func TestInitDBConn(t *testing.T) {
	sqlOpenOri := sqlOpen
	defer func() {
		sqlOpen = sqlOpenOri
	}()

	type args struct {
		cfg configuration.DatabaseConfig
	}
	tests := []struct {
		name        string
		args        args
		mockSqlOpen func(driverName, dataSourceName string) (*sql.DB, error)
		wantErr     error
	}{
		{
			name: "when_sqlOpen_error_then_return_error",
			args: args{},
			mockSqlOpen: func(driverName, dataSourceName string) (*sql.DB, error) {
				return nil, assert.AnError
			},
			wantErr: assert.AnError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sqlOpen = test.mockSqlOpen
			err := InitDBConn(test.args.cfg)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
