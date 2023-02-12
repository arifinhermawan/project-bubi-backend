// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	account "github.com/arifinhermawan/bubi/internal/service/account"
	gomock "github.com/golang/mock/gomock"
)

// MockaccountServiceProvider is a mock of accountServiceProvider interface.
type MockaccountServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockaccountServiceProviderMockRecorder
}

// MockaccountServiceProviderMockRecorder is the mock recorder for MockaccountServiceProvider.
type MockaccountServiceProviderMockRecorder struct {
	mock *MockaccountServiceProvider
}

// NewMockaccountServiceProvider creates a new mock instance.
func NewMockaccountServiceProvider(ctrl *gomock.Controller) *MockaccountServiceProvider {
	mock := &MockaccountServiceProvider{ctrl: ctrl}
	mock.recorder = &MockaccountServiceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockaccountServiceProvider) EXPECT() *MockaccountServiceProviderMockRecorder {
	return m.recorder
}

// CheckPasswordCorrect mocks base method.
func (m *MockaccountServiceProvider) CheckPasswordCorrect(ctx context.Context, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordCorrect", ctx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPasswordCorrect indicates an expected call of CheckPasswordCorrect.
func (mr *MockaccountServiceProviderMockRecorder) CheckPasswordCorrect(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordCorrect", reflect.TypeOf((*MockaccountServiceProvider)(nil).CheckPasswordCorrect), ctx, email, password)
}

// GenerateJWT mocks base method.
func (m *MockaccountServiceProvider) GenerateJWT(ctx context.Context, userID int64, email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWT", ctx, userID, email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWT indicates an expected call of GenerateJWT.
func (mr *MockaccountServiceProviderMockRecorder) GenerateJWT(ctx, userID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWT", reflect.TypeOf((*MockaccountServiceProvider)(nil).GenerateJWT), ctx, userID, email)
}

// GetUserAccountByEmail mocks base method.
func (m *MockaccountServiceProvider) GetUserAccountByEmail(ctx context.Context, email string) (account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAccountByEmail", ctx, email)
	ret0, _ := ret[0].(account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccountByEmail indicates an expected call of GetUserAccountByEmail.
func (mr *MockaccountServiceProviderMockRecorder) GetUserAccountByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccountByEmail", reflect.TypeOf((*MockaccountServiceProvider)(nil).GetUserAccountByEmail), ctx, email)
}

// InsertUserAccount mocks base method.
func (m *MockaccountServiceProvider) InsertUserAccount(ctx context.Context, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserAccount", ctx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserAccount indicates an expected call of InsertUserAccount.
func (mr *MockaccountServiceProviderMockRecorder) InsertUserAccount(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserAccount", reflect.TypeOf((*MockaccountServiceProvider)(nil).InsertUserAccount), ctx, email, password)
}

// InvalidateJWT mocks base method.
func (m *MockaccountServiceProvider) InvalidateJWT(ctx context.Context, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateJWT", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateJWT indicates an expected call of InvalidateJWT.
func (mr *MockaccountServiceProviderMockRecorder) InvalidateJWT(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateJWT", reflect.TypeOf((*MockaccountServiceProvider)(nil).InvalidateJWT), ctx, userID)
}

// UpdateUserAccount mocks base method.
func (m *MockaccountServiceProvider) UpdateUserAccount(ctx context.Context, param account.UpdateUserAccountParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAccount", ctx, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAccount indicates an expected call of UpdateUserAccount.
func (mr *MockaccountServiceProviderMockRecorder) UpdateUserAccount(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAccount", reflect.TypeOf((*MockaccountServiceProvider)(nil).UpdateUserAccount), ctx, param)
}

// UpdateUserPassword mocks base method.
func (m *MockaccountServiceProvider) UpdateUserPassword(ctx context.Context, userID int64, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, userID, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockaccountServiceProviderMockRecorder) UpdateUserPassword(ctx, userID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockaccountServiceProvider)(nil).UpdateUserPassword), ctx, userID, password)
}
