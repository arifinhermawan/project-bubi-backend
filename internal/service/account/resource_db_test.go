package account

import (
	// golang package
	"context"
	"database/sql"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

func TestResource_GetUserAccountByEmailFromDB(t *testing.T) {
	email := "lee.jieun@iu.com"

	type mockFields struct {
		db *MockdbRepoProvider
	}
	tests := []struct {
		name       string
		args       string
		mockFields func(mockFields)
		want       entity.Account
		wantErr    error
	}{
		{
			name: "when_GetUserAccountByEmail_error_then_return_error",
			args: email,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().GetUserAccountByEmail(context.Background(), email).Return(pgsql.Account{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_GetUserAccountByEmail_return_empty_account_then_return_empty_struct",
			args: email,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().GetUserAccountByEmail(context.Background(), email).Return(pgsql.Account{}, nil)
			},
		},
		{
			name: "when_no_error_occured_then_return_populated_account_struct",
			args: email,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().GetUserAccountByEmail(context.Background(), email).Return(pgsql.Account{
					Email:             "lee.jieun@iu.com",
					FirstName:         sql.NullString{Valid: true, String: "Ji Eun"},
					ID:                1,
					LastName:          sql.NullString{Valid: true, String: "Lee"},
					Password:          "password 123",
					RecordPeriodStart: 1,
				}, nil)
			},
			want: entity.Account{
				Email:             "lee.jieun@iu.com",
				FirstName:         "Ji Eun",
				ID:                1,
				LastName:          "Lee",
				Password:          "password 123",
				RecordPeriodStart: 1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				db: NewMockdbRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			rsc := Resource{
				db: mockFields.db,
			}

			got, err := rsc.GetUserAccountByEmailFromDB(context.Background(), test.args)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestResource_InsertUserAccountToDB(t *testing.T) {
	type mockFields struct {
		db *MockdbRepoProvider
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
			name: "when_BeginTX_error_then_return_error",
			args: args{},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_InsertUserAccount_error_then_rollback_transaction_then_return_error",
			args: args{
				email:    "email",
				password: "password",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().InsertUserAccount(context.Background(), &sql.Tx{}, "email", "password").Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(nil)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_InsertUserAccount_error_and_failed_to_rollback_transaction_then_log_the_error_then_return_error",
			args: args{
				email:    "email",
				password: "password",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().InsertUserAccount(context.Background(), &sql.Tx{}, "email", "password").Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_failed_to_commit_then_log_the_error",
			args: args{
				email:    "email",
				password: "password",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().InsertUserAccount(context.Background(), &sql.Tx{}, "email", "password").Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(assert.AnError)
			},
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: args{
				email:    "email",
				password: "password",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().InsertUserAccount(context.Background(), &sql.Tx{}, "email", "password").Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				db: NewMockdbRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			rsc := Resource{
				db: mockFields.db,
			}

			err := rsc.InsertUserAccountToDB(context.Background(), test.args.email, test.args.password)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestResource_UpdateUserAccountInDB(t *testing.T) {
	type mockFields struct {
		db *MockdbRepoProvider
	}

	mockArgs := UpdateUserAccountParam{
		FirstName:    "Ji Eun",
		LastName:     "Lee",
		RecordPeriod: 25,
		UserID:       123,
	}
	tests := []struct {
		name       string
		args       UpdateUserAccountParam
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_BeginTX_error_then_return_error",
			args: UpdateUserAccountParam{},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_UpdateUserAccount_error_then_rollback_transaction_then_return_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserAccount(context.Background(), &sql.Tx{}, pgsql.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				}).Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(nil)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_UpdateUserAccount_error_and_failed_to_rollback_transaction_then_log_the_error_then_return_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserAccount(context.Background(), &sql.Tx{}, pgsql.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				}).Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_failed_to_commit_then_log_the_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserAccount(context.Background(), &sql.Tx{}, pgsql.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				}).Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(assert.AnError)
			},
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserAccount(context.Background(), &sql.Tx{}, pgsql.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 25,
					UserID:       123,
				}).Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				db: NewMockdbRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			rsc := Resource{
				db: mockFields.db,
			}

			err := rsc.UpdateUserAccountInDB(context.Background(), test.args)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestResource_UpdateUserPasswordInDB(t *testing.T) {
	type mockFields struct {
		db *MockdbRepoProvider
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
			name: "when_BeginTX_error_then_return_error",
			args: args{},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_UpdateUserPassword_error_then_rollback_transaction_then_return_error",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserPassword(context.Background(), &sql.Tx{}, int64(123), "pass").Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(nil)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_UpdateUserPassword_error_and_failed_to_rollback_transaction_then_log_the_error_then_return_error",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserPassword(context.Background(), &sql.Tx{}, int64(123), "pass").Return(assert.AnError)
				mf.db.EXPECT().Rollback(&sql.Tx{}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_failed_to_commit_then_log_the_error",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserPassword(context.Background(), &sql.Tx{}, int64(123), "pass").Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(assert.AnError)
			},
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: args{
				userID:   123,
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.db.EXPECT().BeginTX(context.Background(), nil).Return(&sql.Tx{}, nil)
				mf.db.EXPECT().UpdateUserPassword(context.Background(), &sql.Tx{}, int64(123), "pass").Return(nil)
				mf.db.EXPECT().Commit(&sql.Tx{}).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				db: NewMockdbRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			rsc := Resource{
				db: mockFields.db,
			}

			err := rsc.UpdateUserPasswordInDB(context.Background(), test.args.userID, test.args.password)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
