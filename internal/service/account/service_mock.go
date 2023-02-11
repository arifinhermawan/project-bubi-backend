// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"
	time "time"

	entity "github.com/arifinhermawan/bubi/internal/entity"
	configuration "github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
	gomock "github.com/golang/mock/gomock"
)

// MockresourceProvider is a mock of resourceProvider interface.
type MockresourceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockresourceProviderMockRecorder
}

// MockresourceProviderMockRecorder is the mock recorder for MockresourceProvider.
type MockresourceProviderMockRecorder struct {
	mock *MockresourceProvider
}

// NewMockresourceProvider creates a new mock instance.
func NewMockresourceProvider(ctrl *gomock.Controller) *MockresourceProvider {
	mock := &MockresourceProvider{ctrl: ctrl}
	mock.recorder = &MockresourceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockresourceProvider) EXPECT() *MockresourceProviderMockRecorder {
	return m.recorder
}

// DeleteJWTInCache mocks base method.
func (m *MockresourceProvider) DeleteJWTInCache(ctx context.Context, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJWTInCache", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJWTInCache indicates an expected call of DeleteJWTInCache.
func (mr *MockresourceProviderMockRecorder) DeleteJWTInCache(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJWTInCache", reflect.TypeOf((*MockresourceProvider)(nil).DeleteJWTInCache), ctx, userID)
}

// GetJWTFromCache mocks base method.
func (m *MockresourceProvider) GetJWTFromCache(ctx context.Context, userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTFromCache", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJWTFromCache indicates an expected call of GetJWTFromCache.
func (mr *MockresourceProviderMockRecorder) GetJWTFromCache(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTFromCache", reflect.TypeOf((*MockresourceProvider)(nil).GetJWTFromCache), ctx, userID)
}

// GetUserAccountByEmailFromDB mocks base method.
func (m *MockresourceProvider) GetUserAccountByEmailFromDB(ctx context.Context, email string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAccountByEmailFromDB", ctx, email)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccountByEmailFromDB indicates an expected call of GetUserAccountByEmailFromDB.
func (mr *MockresourceProviderMockRecorder) GetUserAccountByEmailFromDB(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccountByEmailFromDB", reflect.TypeOf((*MockresourceProvider)(nil).GetUserAccountByEmailFromDB), ctx, email)
}

// InsertUserAccountToDB mocks base method.
func (m *MockresourceProvider) InsertUserAccountToDB(ctx context.Context, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserAccountToDB", ctx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserAccountToDB indicates an expected call of InsertUserAccountToDB.
func (mr *MockresourceProviderMockRecorder) InsertUserAccountToDB(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserAccountToDB", reflect.TypeOf((*MockresourceProvider)(nil).InsertUserAccountToDB), ctx, email, password)
}

// SetJWTToCache mocks base method.
func (m *MockresourceProvider) SetJWTToCache(ctx context.Context, userID int64, jwt string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetJWTToCache", ctx, userID, jwt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJWTToCache indicates an expected call of SetJWTToCache.
func (mr *MockresourceProviderMockRecorder) SetJWTToCache(ctx, userID, jwt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJWTToCache", reflect.TypeOf((*MockresourceProvider)(nil).SetJWTToCache), ctx, userID, jwt)
}

// UpdateUserAccountInDB mocks base method.
func (m *MockresourceProvider) UpdateUserAccountInDB(ctx context.Context, param UpdateUserAccountParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAccountInDB", ctx, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAccountInDB indicates an expected call of UpdateUserAccountInDB.
func (mr *MockresourceProviderMockRecorder) UpdateUserAccountInDB(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAccountInDB", reflect.TypeOf((*MockresourceProvider)(nil).UpdateUserAccountInDB), ctx, param)
}

// MockinfraProvider is a mock of infraProvider interface.
type MockinfraProvider struct {
	ctrl     *gomock.Controller
	recorder *MockinfraProviderMockRecorder
}

// MockinfraProviderMockRecorder is the mock recorder for MockinfraProvider.
type MockinfraProviderMockRecorder struct {
	mock *MockinfraProvider
}

// NewMockinfraProvider creates a new mock instance.
func NewMockinfraProvider(ctrl *gomock.Controller) *MockinfraProvider {
	mock := &MockinfraProvider{ctrl: ctrl}
	mock.recorder = &MockinfraProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinfraProvider) EXPECT() *MockinfraProviderMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockinfraProvider) GetConfig() *configuration.AppConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*configuration.AppConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockinfraProviderMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockinfraProvider)(nil).GetConfig))
}

// GetTimeGMT7 mocks base method.
func (m *MockinfraProvider) GetTimeGMT7() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeGMT7")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeGMT7 indicates an expected call of GetTimeGMT7.
func (mr *MockinfraProviderMockRecorder) GetTimeGMT7() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeGMT7", reflect.TypeOf((*MockinfraProvider)(nil).GetTimeGMT7))
}
