package account

import (
	// golang package
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HandleUserSignUp(t *testing.T) {
	type mockFields struct {
		accountUC *MockaccountUCManager
		infra     *MockinfraProvider
	}

	tests := []struct {
		name       string
		mockFields func(mockFields)
	}{
		{
			name: "when_ReadAll_error_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, assert.AnError)
			},
		},
		{
			name: "when_JsonUnmarshal_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var dest userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(nil, &dest).Return(assert.AnError)
			},
		},
		{
			name: "when_email_is_empty_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mockParam := userSignUpParam{
					Password: "1234",
				}
				bytesParam, _ := json.Marshal(mockParam)

				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(bytesParam, nil)

				var destination userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(bytesParam, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*userSignUpParam) = userSignUpParam{
							Password: "1234",
						}
						return nil
					})
			},
		},
		{
			name: "when_password_is_empty_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mockParam := userSignUpParam{
					Email: "email",
				}
				bytesParam, _ := json.Marshal(mockParam)
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(bytesParam, nil)

				var destination userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(bytesParam, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*userSignUpParam) = userSignUpParam{
							Email: "email",
						}
						return nil
					})
			},
		},
		{
			name: "when_UserSignUp_error_then_return_internal_server_error",
			mockFields: func(mf mockFields) {
				mockParam := userSignUpParam{
					Email:    "email",
					Password: "password",
				}
				bytesParam, _ := json.Marshal(mockParam)
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(bytesParam, nil)

				var destination userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(bytesParam, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*userSignUpParam) = userSignUpParam{
							Email:    "email",
							Password: "password",
						}
						return nil
					})

				mf.accountUC.EXPECT().UserSignUp(context.Background(), "email", "password").Return(assert.AnError)
			},
		},
		{
			name: "when_user_already_exist_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mockParam := userSignUpParam{
					Email:    "email",
					Password: "password",
				}
				bytesParam, _ := json.Marshal(mockParam)
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(bytesParam, nil)

				var destination userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(bytesParam, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*userSignUpParam) = userSignUpParam{
							Email:    "email",
							Password: "password",
						}
						return nil
					})

				mf.accountUC.EXPECT().UserSignUp(context.Background(), "email", "password").Return(errUserExist)
			},
		},
		{
			name: "when_no_error_occured_then_return_status_created",
			mockFields: func(mf mockFields) {
				mockParam := userSignUpParam{
					Email:    "email",
					Password: "password",
				}
				bytesParam, _ := json.Marshal(mockParam)
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(bytesParam, nil)

				var destination userSignUpParam
				mf.infra.EXPECT().JsonUnmarshal(bytesParam, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*userSignUpParam) = userSignUpParam{
							Email:    "email",
							Password: "password",
						}
						return nil
					})

				mf.accountUC.EXPECT().UserSignUp(context.Background(), "email", "password").Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				accountUC: NewMockaccountUCManager(ctrl),
				infra:     NewMockinfraProvider(ctrl),
			}
			test.mockFields(mockFields)

			h := &Handler{
				account: mockFields.accountUC,
				infra:   mockFields.infra,
			}

			req := httptest.NewRequest(http.MethodPost, "/account/signup", nil)
			w := httptest.NewRecorder()

			h.HandleUserSignUp(w, req)
		})
	}
}
