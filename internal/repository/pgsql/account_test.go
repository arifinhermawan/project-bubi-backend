package pgsql

import (
	// golang package
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	// external package
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

var (
	mockConfig = &configuration.AppConfig{
		Database: configuration.DatabaseConfig{
			DefaultTimeout: 5,
		},
	}
)

func TestDBRepository_GetUserAccountByEmail(t *testing.T) {
	funcSQLXNamedOri := sqlx.Named

	expectedQuery := `
		SELECT
			email,
			record_period_start,
			first_name,
			last_name,
			id,
			password
		FROM
			user_account
		WHERE
			email = $1
	`

	type mockFields struct {
		infra *MockinfraProvider
		sql   sqlmock.Sqlmock
	}

	type args struct {
		email string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		want       Account
		wantErr    error
	}{
		{
			name: "when_funcSQLXNamed_error_then_return_error",
			args: args{
				email: "lee.jieun@iu.com",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)

				funcSQLXNamed = func(query string, arg interface{}) (string, []interface{}, error) {
					return "", nil, assert.AnError
				}
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_GetContext_error_then_return_empty_struct_and_an_error",
			args: args{
				email: "lee.jieun@iu.com",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.sql.ExpectQuery(expectedQuery).WillReturnError(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_and_account_not_exist_then_return_empty_account",
			args: args{
				email: "lee.jieun@iu.com",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.sql.ExpectQuery(expectedQuery).WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
		},
		{
			name: "when_no_error_occured_and_account_exist_then_return_the_account",
			args: args{
				email: "lee.jieun@iu.com",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)

				rows := sqlmock.NewRows([]string{"email", "first_name", "id", "last_name", "password", "record_period_start"}).
					AddRow(
						"lee.jieun@iu.com",
						"Ji Eun",
						"1",
						"Lee",
						"ijigeum",
						"1",
					)
				mf.sql.ExpectQuery(expectedQuery).WithArgs("lee.jieun@iu.com").WillReturnRows(rows)
			},
			want: Account{
				Email:             "lee.jieun@iu.com",
				FirstName:         sql.NullString{String: "Ji Eun", Valid: true},
				ID:                1,
				LastName:          sql.NullString{String: "Lee", Valid: true},
				Password:          "ijigeum",
				RecordPeriodStart: 1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				assert.Nil(t, err)
				return
			}

			defer func() {
				mockDB.Close()
				funcSQLXNamed = funcSQLXNamedOri
			}()

			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				infra: NewMockinfraProvider(ctrl),
				sql:   mockSql,
			}
			test.mockFields(mockFields)

			r := DBRepository{
				infra: mockFields.infra,
				db:    sqlx.NewDb(mockDB, "postgres"),
			}

			got, err := r.GetUserAccountByEmail(context.Background(), test.want.Email)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}

func TestDBRepository_InsertUserAccount(t *testing.T) {
	funcSQLXNamedOri := sqlx.Named
	mockTime := time.Date(1993, 05, 16, 0, 0, 0, 0, time.UTC)

	expectedQuery := `
		INSERT INTO 
			user_account(email,"password",created_at)
		VALUES (
			$1,
			$2,
			$3
		)
	`

	type mockFields struct {
		infra *MockinfraProvider
		sql   sqlmock.Sqlmock
	}

	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_funcSQLXNamed_error_then_return_error",
			args: args{
				email:    "lee.jieun@iu.com",
				password: "abcd",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				funcSQLXNamed = func(query string, arg interface{}) (string, []interface{}, error) {
					return "", nil, assert.AnError
				}
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_ExecContext_error_then_return_error",
			args: args{
				email:    "lee.jieun@iu.com",
				password: "abcd",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				mf.sql.ExpectExec(expectedQuery).WillReturnError(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				email:    "lee.jieun@iu.com",
				password: "abcd",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				mf.sql.ExpectExec(expectedQuery).WithArgs("lee.jieun@iu.com", "abcd", mockTime).WillReturnResult(sqlmock.NewErrorResult(nil))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				assert.Nil(t, err)
				return
			}

			defer func() {
				mockDB.Close()
				funcSQLXNamed = funcSQLXNamedOri
			}()

			mockSQL.ExpectBegin().WillReturnError(nil)
			tx, _ := mockDB.Begin()

			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				infra: NewMockinfraProvider(ctrl),
				sql:   mockSQL,
			}
			test.mockFields(mockFields)

			r := DBRepository{
				infra: mockFields.infra,
				db:    sqlx.NewDb(mockDB, "postgres"),
			}

			err = r.InsertUserAccount(context.Background(), tx, test.args.email, test.args.password)
			assert.Equal(t, test.wantErr, err)
			assert.Nil(t, mockSQL.ExpectationsWereMet())
		})
	}
}

func TestDBRepository_UpdateUserAccount(t *testing.T) {
	funcSQLXNamedOri := sqlx.Named
	mockTime := time.Date(1993, 05, 16, 0, 0, 0, 0, time.UTC)

	expectedQuery := `
		UPDATE
			user_account
		SET 
			first_name = $1,
			last_name = $2,
			record_period_start = $3,
			updated_at = $4
		WHERE
			id = $5
	`

	type mockFields struct {
		infra *MockinfraProvider
		sql   sqlmock.Sqlmock
	}

	type args struct {
		param UpdateUserAccountParam
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_funcSQLXNamed_error_then_return_error",
			args: args{},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				funcSQLXNamed = func(query string, arg interface{}) (string, []interface{}, error) {
					return "", nil, assert.AnError
				}
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_ExecContext_error_then_return_error",
			args: args{
				param: UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				},
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				mf.sql.ExpectExec(expectedQuery).WillReturnError(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				param: UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				},
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.infra.EXPECT().GetTimeGMT7().Return(mockTime)

				mf.sql.ExpectExec(expectedQuery).
					WithArgs(
						"Ji Eun",
						"Lee",
						25,
						mockTime,
						int64(123),
					).WillReturnResult(driver.RowsAffected(1))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				assert.Nil(t, err)
				return
			}

			defer func() {
				mockDB.Close()
				funcSQLXNamed = funcSQLXNamedOri
			}()

			mockSQL.ExpectBegin().WillReturnError(nil)
			tx, _ := mockDB.Begin()

			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				infra: NewMockinfraProvider(ctrl),
				sql:   mockSQL,
			}
			test.mockFields(mockFields)

			r := DBRepository{
				infra: mockFields.infra,
				db:    sqlx.NewDb(mockDB, "postgres"),
			}

			err = r.UpdateUserAccount(context.Background(), tx, test.args.param)
			assert.Equal(t, test.wantErr, err)
			assert.Nil(t, mockSQL.ExpectationsWereMet())
		})
	}
}

func TestDBRepository_UpdateUserPassword(t *testing.T) {
	funcSQLXNamedOri := sqlx.Named

	expectedQuery := `
		UPDATE
			user_account
		SET
			password = $1
		WHERE
			id = $2
	`

	type mockFields struct {
		infra *MockinfraProvider
		sql   sqlmock.Sqlmock
	}

	type args struct {
		userID   int64
		password string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_funcSQLXNamed_error_then_return_error",
			args: args{},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)

				funcSQLXNamed = func(query string, arg interface{}) (string, []interface{}, error) {
					return "", nil, assert.AnError
				}
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_ExecContext_error_then_return_error",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)

				mf.sql.ExpectExec(expectedQuery).WillReturnError(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)

				mf.sql.ExpectExec(expectedQuery).
					WithArgs(
						"pass",
						int64(123),
					).WillReturnResult(driver.RowsAffected(1))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				assert.Nil(t, err)
				return
			}

			defer func() {
				mockDB.Close()
				funcSQLXNamed = funcSQLXNamedOri
			}()

			mockSQL.ExpectBegin().WillReturnError(nil)
			tx, _ := mockDB.Begin()

			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				infra: NewMockinfraProvider(ctrl),
				sql:   mockSQL,
			}
			test.mockFields(mockFields)

			r := DBRepository{
				infra: mockFields.infra,
				db:    sqlx.NewDb(mockDB, "postgres"),
			}

			err = r.UpdateUserPassword(context.Background(), tx, test.args.userID, test.args.password)
			assert.Equal(t, test.wantErr, err)
			assert.Nil(t, mockSQL.ExpectationsWereMet())
		})
	}
}
