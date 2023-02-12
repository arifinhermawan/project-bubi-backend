package account

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

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

func TestService_UpdateUserAccount(t *testing.T) {
	type mockFields struct {
		rsc *MockresourceProvider
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
			name: "when_UpdateUserAccountInDB_error_then_return_error",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().UpdateUserAccountInDB(context.Background(), mockArgs).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: mockArgs,
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().UpdateUserAccountInDB(context.Background(), mockArgs).Return(nil)
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

			err := svc.UpdateUserAccount(context.Background(), test.args)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestService_UpdateUserPassword(t *testing.T) {
	generateFromPasswordOri := bcrypt.GenerateFromPassword
	type mockFields struct {
		rsc *MockresourceProvider
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
			name: "when_UpdateUserPasswordInDB_error_then_return_error",
			args: args{
				password: "1234",
				userID:   123,
			},
			mockFields: func(mf mockFields) {
				generateFromPassword = func(password []byte, cost int) ([]byte, error) {
					hashed, _ := json.Marshal("1234")
					return hashed, nil
				}

				hashed, _ := json.Marshal("1234")
				mf.rsc.EXPECT().UpdateUserPasswordInDB(context.Background(), int64(123), string(hashed)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: args{
				password: "1234",
				userID:   123,
			},
			mockFields: func(mf mockFields) {
				generateFromPassword = func(password []byte, cost int) ([]byte, error) {
					hashed, _ := json.Marshal("1234")
					return hashed, nil
				}

				hashed, _ := json.Marshal("1234")
				mf.rsc.EXPECT().UpdateUserPasswordInDB(context.Background(), int64(123), string(hashed)).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				generateFromPassword = generateFromPasswordOri
			}()

			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				rsc: NewMockresourceProvider(ctrl),
			}
			test.mockFields(mockFields)

			svc := &Service{
				rsc: mockFields.rsc,
			}

			err := svc.UpdateUserPassword(context.Background(), test.args.userID, test.args.password)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
