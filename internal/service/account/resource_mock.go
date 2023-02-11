// Code generated by MockGen. DO NOT EDIT.
// Source: ./resource.go

// Package account is a generated GoMock package.
package account

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	time "time"

	configuration "github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
	pgsql "github.com/arifinhermawan/bubi/internal/repository/pgsql"
	gomock "github.com/golang/mock/gomock"
)

// MockdbRepoProvider is a mock of dbRepoProvider interface.
type MockdbRepoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockdbRepoProviderMockRecorder
}

// MockdbRepoProviderMockRecorder is the mock recorder for MockdbRepoProvider.
type MockdbRepoProviderMockRecorder struct {
	mock *MockdbRepoProvider
}

// NewMockdbRepoProvider creates a new mock instance.
func NewMockdbRepoProvider(ctrl *gomock.Controller) *MockdbRepoProvider {
	mock := &MockdbRepoProvider{ctrl: ctrl}
	mock.recorder = &MockdbRepoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdbRepoProvider) EXPECT() *MockdbRepoProviderMockRecorder {
	return m.recorder
}

// BeginTX mocks base method.
func (m *MockdbRepoProvider) BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTX", ctx, options)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTX indicates an expected call of BeginTX.
func (mr *MockdbRepoProviderMockRecorder) BeginTX(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTX", reflect.TypeOf((*MockdbRepoProvider)(nil).BeginTX), ctx, options)
}

// Commit mocks base method.
func (m *MockdbRepoProvider) Commit(tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockdbRepoProviderMockRecorder) Commit(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockdbRepoProvider)(nil).Commit), tx)
}

// GetUserAccountByEmail mocks base method.
func (m *MockdbRepoProvider) GetUserAccountByEmail(ctx context.Context, email string) (pgsql.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAccountByEmail", ctx, email)
	ret0, _ := ret[0].(pgsql.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccountByEmail indicates an expected call of GetUserAccountByEmail.
func (mr *MockdbRepoProviderMockRecorder) GetUserAccountByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccountByEmail", reflect.TypeOf((*MockdbRepoProvider)(nil).GetUserAccountByEmail), ctx, email)
}

// InsertUserAccount mocks base method.
func (m *MockdbRepoProvider) InsertUserAccount(ctx context.Context, tx *sql.Tx, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserAccount", ctx, tx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserAccount indicates an expected call of InsertUserAccount.
func (mr *MockdbRepoProviderMockRecorder) InsertUserAccount(ctx, tx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserAccount", reflect.TypeOf((*MockdbRepoProvider)(nil).InsertUserAccount), ctx, tx, email, password)
}

// Rollback mocks base method.
func (m *MockdbRepoProvider) Rollback(tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockdbRepoProviderMockRecorder) Rollback(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockdbRepoProvider)(nil).Rollback), tx)
}

// UpdateUserAccount mocks base method.
func (m *MockdbRepoProvider) UpdateUserAccount(ctx context.Context, tx *sql.Tx, param pgsql.UpdateUserAccountParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAccount", ctx, tx, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAccount indicates an expected call of UpdateUserAccount.
func (mr *MockdbRepoProviderMockRecorder) UpdateUserAccount(ctx, tx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAccount", reflect.TypeOf((*MockdbRepoProvider)(nil).UpdateUserAccount), ctx, tx, param)
}

// MockinfraRepoProvider is a mock of infraRepoProvider interface.
type MockinfraRepoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockinfraRepoProviderMockRecorder
}

// MockinfraRepoProviderMockRecorder is the mock recorder for MockinfraRepoProvider.
type MockinfraRepoProviderMockRecorder struct {
	mock *MockinfraRepoProvider
}

// NewMockinfraRepoProvider creates a new mock instance.
func NewMockinfraRepoProvider(ctrl *gomock.Controller) *MockinfraRepoProvider {
	mock := &MockinfraRepoProvider{ctrl: ctrl}
	mock.recorder = &MockinfraRepoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinfraRepoProvider) EXPECT() *MockinfraRepoProviderMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockinfraRepoProvider) GetConfig() *configuration.AppConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*configuration.AppConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockinfraRepoProviderMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockinfraRepoProvider)(nil).GetConfig))
}

// JsonUnmarshal mocks base method.
func (m *MockinfraRepoProvider) JsonUnmarshal(input []byte, dest interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JsonUnmarshal", input, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// JsonUnmarshal indicates an expected call of JsonUnmarshal.
func (mr *MockinfraRepoProviderMockRecorder) JsonUnmarshal(input, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JsonUnmarshal", reflect.TypeOf((*MockinfraRepoProvider)(nil).JsonUnmarshal), input, dest)
}

// MockredisRepoProvider is a mock of redisRepoProvider interface.
type MockredisRepoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockredisRepoProviderMockRecorder
}

// MockredisRepoProviderMockRecorder is the mock recorder for MockredisRepoProvider.
type MockredisRepoProviderMockRecorder struct {
	mock *MockredisRepoProvider
}

// NewMockredisRepoProvider creates a new mock instance.
func NewMockredisRepoProvider(ctrl *gomock.Controller) *MockredisRepoProvider {
	mock := &MockredisRepoProvider{ctrl: ctrl}
	mock.recorder = &MockredisRepoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockredisRepoProvider) EXPECT() *MockredisRepoProviderMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockredisRepoProvider) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockredisRepoProviderMockRecorder) Del(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockredisRepoProvider)(nil).Del), ctx, key)
}

// Get mocks base method.
func (m *MockredisRepoProvider) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockredisRepoProviderMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockredisRepoProvider)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockredisRepoProvider) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockredisRepoProviderMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockredisRepoProvider)(nil).Set), ctx, key, value, expiration)
}
