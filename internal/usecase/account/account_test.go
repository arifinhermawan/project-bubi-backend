package account

import (
	// golang package
	"context"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
			name: "when_CheckIsAccountExist_error_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().CheckIsAccountExist(context.Background(), "email").Return(false, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_CheckIsAccountExist_return_true_then_return_error",
			args: args{email: "email"},
			mockFields: func(mf mockFields) {
				mf.accountSvc.EXPECT().CheckIsAccountExist(context.Background(), "email").Return(true, nil)
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
				mf.accountSvc.EXPECT().CheckIsAccountExist(context.Background(), "email").Return(false, nil)
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
				mf.accountSvc.EXPECT().CheckIsAccountExist(context.Background(), "email").Return(false, nil)
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
