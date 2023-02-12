package account

import (
	// golang package
	"context"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/service/account"
)

func TestUseCase_LogIn(t *testing.T) {
	type mockFields struct {
		accountSvc *MockaccountServiceProvider
	}

	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		want       string
		wantErr    error
	}{
		{
			name: "when_GetUserAccountByEmail_error_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_account_not_exist_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{}, nil)
			},
			wantErr: errUserNotExist,
		},
		{
			name: "when_CheckPasswordCorrect_error_then_return_error",
			args: args{
				email:    "email",
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{
					ID: 123,
				}, nil)
				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "pass").Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_GenerateJWT_error_then_return_error",
			args: args{
				email:    "email",
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{
					ID: 123,
				}, nil)

				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "pass").Return(nil)
				mf.accountSvc.EXPECT().GenerateJWT(context.Background(), int64(123), "email").Return("", assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				email:    "email",
				password: "pass",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{
					ID: 123,
				}, nil)

				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "pass").Return(nil)
				mf.accountSvc.EXPECT().GenerateJWT(context.Background(), int64(123), "email").Return("abc", nil)
			},
			want: "abc",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountSvc: NewMockaccountServiceProvider(ctrl),
			}
			test.mockFields(mockFields)

			uc := &UseCase{
				account: mockFields.accountSvc,
			}

			got, err := uc.LogIn(context.Background(), test.args.email, test.args.password)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestUseCase_LogOut(t *testing.T) {
	type mockFields struct {
		accountSvc *MockaccountServiceProvider
	}

	type args struct {
		userID int64
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_InvalidateJWT_then_return_error",
			args: args{userID: 1234},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().InvalidateJWT(context.Background(), int64(1234)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{userID: 1234},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().InvalidateJWT(context.Background(), int64(1234)).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountSvc: NewMockaccountServiceProvider(ctrl),
			}
			test.mockFields(mockFields)

			uc := &UseCase{
				account: mockFields.accountSvc,
			}

			err := uc.LogOut(context.Background(), test.args.userID)
			assert.Equal(t, test.wantErr, err)

		})
	}
}

func TestUseCase_UpdateUserAccount(t *testing.T) {
	type mockFields struct {
		accountSvc *MockaccountServiceProvider
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
			name: "when_UpdateUserAccount_error_then_return_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().UpdateUserAccount(context.Background(), account.UpdateUserAccountParam(mockArgs)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().UpdateUserAccount(context.Background(), account.UpdateUserAccountParam(mockArgs)).Return(nil)
			},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountSvc: NewMockaccountServiceProvider(ctrl),
			}
			test.mockFields(mockFields)

			uc := &UseCase{
				account: mockFields.accountSvc,
			}

			err := uc.UpdateUserAccount(context.Background(), test.args)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestUseCase_UserSignUp(t *testing.T) {
	type mockFields struct {
		accountSvc *MockaccountServiceProvider
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
			name: "when_GetUserAccountByEmail_error_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_account_exist_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{ID: 123}, nil)
			},
			wantErr: errUserExist,
		},
		{
			name: "when_InsertUserAccount_error_then_return_error",
			args: args{
				email:    "email",
				password: "passw0rd",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{}, nil)
				mf.accountSvc.EXPECT().InsertUserAccount(context.Background(), "email", "passw0rd").Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: args{
				email:    "email",
				password: "passw0rd",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().GetUserAccountByEmail(context.Background(), "email").Return(account.Account{}, nil)
				mf.accountSvc.EXPECT().InsertUserAccount(context.Background(), "email", "passw0rd").Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountSvc: NewMockaccountServiceProvider(ctrl),
			}
			test.mockFields(mockFields)

			uc := &UseCase{
				account: mockFields.accountSvc,
			}

			err := uc.UserSignUp(context.Background(), test.args.email, test.args.password)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestUseCase_UpdatePassword(t *testing.T) {
	type mockFields struct {
		accountSvc *MockaccountServiceProvider
	}
	tests := []struct {
		name       string
		args       UpdatePasswordParam
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_CheckPasswordCorrect_error_then_return_error",
			args: UpdatePasswordParam{
				Email:       "email",
				OldPassword: "oldpass",
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "oldpass").Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_UpdateUserPassword_error_then_return_error",
			args: UpdatePasswordParam{
				Email:       "email",
				OldPassword: "oldpass",
				Password:    "password",
				UserID:      123,
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "oldpass").Return(nil)
				mf.accountSvc.EXPECT().UpdateUserPassword(context.Background(), int64(123), "password").Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: UpdatePasswordParam{
				Email:       "email",
				OldPassword: "oldpass",
				Password:    "password",
				UserID:      123,
			},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().CheckPasswordCorrect(context.Background(), "email", "oldpass").Return(nil)
				mf.accountSvc.EXPECT().UpdateUserPassword(context.Background(), int64(123), "password").Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountSvc: NewMockaccountServiceProvider(ctrl),
			}
			test.mockFields(mockFields)

			uc := &UseCase{
				account: mockFields.accountSvc,
			}

			err := uc.UpdatePassword(context.Background(), test.args)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
