package account

import (
	// golang package
	"context"
	"encoding/json"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	// internal package
)

// func TestService_CheckIsAccountExist(t *testing.T) {
// 	mockEmail := "jieun.lee@iu.com"

// 	type mockFields struct {
// 		rsc *MockresourceProvider
// 	}
// 	tests := []struct {
// 		name       string
// 		args       string
// 		mockFields func(mockFields)
// 		want       bool
// 		wantErr    error
// 	}{
// 		{
// 			name: "when_GetUserAccountByEmailFromDB_error_then_return_error",
// 			args: mockEmail,
// 			mockFields: func(mf mockFields) {
// 				mf.rsc.EXPECT().GetUserAccountByEmailFromDB(context.Background(), mockEmail).Return(entity.Account{}, assert.AnError)
// 			},
// 			wantErr: assert.AnError,
// 		},
// 		{
// 			name: "when_account_exist_then_return_true",
// 			args: mockEmail,
// 			mockFields: func(mf mockFields) {
// 				mf.rsc.EXPECT().GetUserAccountByEmailFromDB(context.Background(), mockEmail).Return(entity.Account{ID: 123}, nil)
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "when_account_not_exist_then_return_false",
// 			args: mockEmail,
// 			mockFields: func(mf mockFields) {
// 				mf.rsc.EXPECT().GetUserAccountByEmailFromDB(context.Background(), mockEmail).Return(entity.Account{}, nil)
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			mockFields := mockFields{
// 				rsc: NewMockresourceProvider(ctrl),
// 			}
// 			test.mockFields(mockFields)

// 			svc := &Service{
// 				rsc: mockFields.rsc,
// 			}

// 			got, err := svc.CheckIsAccountExist(context.Background(), test.args)
// 			assert.Equal(t, test.want, got)
// 			assert.Equal(t, test.wantErr, err)
// 		})
// 	}
// }

func TestService_InsertUserAccount(t *testing.T) {
	generateFromPasswordOri := bcrypt.GenerateFromPassword
	type mockFields struct {
		rsc *MockresourceProvider
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
			name: "when_generateFromPassword_error_then_return_error",
			args: args{password: "1234"},
			mockFields: func(mf mockFields) {
				generateFromPassword = func(password []byte, cost int) ([]byte, error) {
					return nil, assert.AnError
				}
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_InsertUserAccountToDB_error_then_return_error",
			args: args{
				password: "1234",
				email:    "email",
			},
			mockFields: func(mf mockFields) {
				generateFromPassword = func(password []byte, cost int) ([]byte, error) {
					hashed, _ := json.Marshal("1234")
					return hashed, nil
				}

				hashed, _ := json.Marshal("1234")
				mf.rsc.EXPECT().InsertUserAccountToDB(context.Background(), "email", string(hashed)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				password: "1234",
				email:    "email",
			},
			mockFields: func(mf mockFields) {
				generateFromPassword = func(password []byte, cost int) ([]byte, error) {
					hashed, _ := json.Marshal("1234")
					return hashed, nil
				}

				hashed, _ := json.Marshal("1234")
				mf.rsc.EXPECT().InsertUserAccountToDB(context.Background(), "email", string(hashed)).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				rsc: NewMockresourceProvider(ctrl),
			}
			test.mockFields(mockFields)

			svc := &Service{
				rsc: mockFields.rsc,
			}

			defer func() {
				generateFromPassword = generateFromPasswordOri
			}()

			err := svc.InsertUserAccount(context.Background(), test.args.email, test.args.password)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
