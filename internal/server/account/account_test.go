package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/arifinhermawan/bubi/internal/usecase/account"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HandleUserLogIn(t *testing.T) {
	type mockFields struct {
		accountUC *MockaccountUCManager
		infra     *MockinfraProvider
	}
	tests := []struct {
		name          string
		emailValid    bool
		passwordValid bool
		mockFields    func(mockFields)
	}{
		{
			name:          "when_email_empty_then_return_bad_request",
			emailValid:    false,
			passwordValid: true,
			mockFields:    func(mf mockFields) {},
		},
		{
			name:       "when_password_empty_then_return_bad_request",
			emailValid: true,
			mockFields: func(mf mockFields) {},
		},
		{
			name:          "when_LogIn_error_then_return_internal_server_error",
			emailValid:    true,
			passwordValid: true,
			mockFields: func(mf mockFields) {
				mf.accountUC.EXPECT().LogIn(context.Background(), "email", "password").Return("", assert.AnError)
			},
		},
		{
			name:          "when_no_error_occured_then_return_status_ok",
			emailValid:    true,
			passwordValid: true,
			mockFields: func(mf mockFields) {
				mf.accountUC.EXPECT().LogIn(context.Background(), "email", "password").Return("", nil)
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

			req := httptest.NewRequest(http.MethodPost, "/account/login", nil)
			req.Form = url.Values{
				"email":    []string{"email"},
				"password": []string{"password"},
			}

			if !test.emailValid {
				req.Form = url.Values{
					"email": []string{""},
				}
			}

			if !test.passwordValid {
				req.Form = url.Values{
					"email":    []string{"email"},
					"password": []string{},
				}
			}

			w := httptest.NewRecorder()

			h.HandleUserLogIn(w, req)
		})
	}
}

func TestHandler_HandlerUserLogOut(t *testing.T) {
	type mockFields struct {
		accountUC *MockaccountUCManager
		infra     *MockinfraProvider
	}
	tests := []struct {
		name        string
		userIDValid bool
		mockFields  func(mockFields)
	}{
		{
			name:        "when_user_id_invalid_then_return_immediately",
			userIDValid: false,
			mockFields:  func(mf mockFields) {},
		},
		{
			name:        "when_LogOut_error_then_return_immediately",
			userIDValid: true,
			mockFields: func(mf mockFields) {
				mf.accountUC.EXPECT().LogOut(context.Background(), int64(1234)).Return(assert.AnError)
			},
		},
		{
			name:        "when_no_error_occured_then_return_success",
			userIDValid: true,
			mockFields: func(mf mockFields) {
				mf.accountUC.EXPECT().LogOut(context.Background(), int64(1234)).Return(nil)
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

			req := httptest.NewRequest(http.MethodPost, "/account/logout", nil)
			req.Form = url.Values{
				"user_id": []string{"abc"},
			}
			if test.userIDValid {
				req.Form = url.Values{
					"user_id": []string{"1234"},
				}
			}

			w := httptest.NewRecorder()

			h.HandlerUserLogOut(w, req)
		})
	}
}

func TestHandler_HandleUpdateUserAccount(t *testing.T) {
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

				var dest updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &dest).Return(assert.AnError)
			},
		},
		{
			name: "when_first_name_is_empty_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var dest updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &dest).Return(nil)
			},
		},
		{
			name: "when_record_period_0_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var destination updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*updateUserAccount) = updateUserAccount{
							FirstName:    "Ji Eun",
							LastName:     "Lee",
							RecordPeriod: 0,
						}

						return nil
					})
			},
		},
		{
			name: "when_user_id_not_valid_then_return_bad_request",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var destination updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*updateUserAccount) = updateUserAccount{
							FirstName:    "Ji Eun",
							LastName:     "Lee",
							RecordPeriod: 1,
							UserID:       0,
						}

						return nil
					})
			},
		},
		{
			name: "when_UpdateUserAccount_error_then_return_internal_server_error",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var destination updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*updateUserAccount) = updateUserAccount{
							FirstName:    "Ji Eun",
							LastName:     "Lee",
							RecordPeriod: 1,
							UserID:       1234,
						}

						return nil
					})

				mf.accountUC.EXPECT().UpdateUserAccount(context.Background(), account.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 1,
					UserID:       1234,
				}).Return(assert.AnError)
			},
		},
		{
			name: "when_no_error_then_return_status_ok",
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().ReadAll(gomock.Any()).Return(nil, nil)

				var destination updateUserAccount
				mf.infra.EXPECT().JsonUnmarshal(nil, &destination).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*updateUserAccount) = updateUserAccount{
							FirstName:    "Ji Eun",
							LastName:     "Lee",
							RecordPeriod: 1,
							UserID:       1234,
						}

						return nil
					})

				mf.accountUC.EXPECT().UpdateUserAccount(context.Background(), account.UpdateUserAccountParam{
					FirstName:    "Ji Eun",
					LastName:     "Lee",
					RecordPeriod: 1,
					UserID:       1234,
				}).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/account/update", nil)

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

			w := httptest.NewRecorder()

			h.HandleUpdateUserAccount(w, req)
		})
	}
}

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
