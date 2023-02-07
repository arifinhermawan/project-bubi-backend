package account

import (
	// golang
	"context"
	"testing"
	"time"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

func TestResource_DeleteJWTInCache(t *testing.T) {
	mockKey := "account:jwt:3"
	type mockFields struct {
		cache *MockredisRepoProvider
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
			name: "name_when_Del_error_then_return_error",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Del(context.Background(), mockKey).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Del(context.Background(), mockKey).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				cache: NewMockredisRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			r := &Resource{
				cache: mockFields.cache,
			}

			err := r.DeleteJWTInCache(context.Background(), test.args.userID)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestResource_GetJWTFromCache(t *testing.T) {
	mockKey := "account:jwt:3"
	type mockFields struct {
		cache *MockredisRepoProvider
		infra *MockinfraRepoProvider
	}

	type args struct {
		userID int64
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		want       string
		wantErr    error
	}{
		{
			name: "when_Get_error_then_return_error",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Get(context.Background(), mockKey).Return("", assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_key_not_exist_then_return_empty_string",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Get(context.Background(), mockKey).Return("", nil)
			},
		},
		{
			name: "when_failed_to_unmarshal_then_return_error",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Get(context.Background(), mockKey).Return("abcd", nil)

				var dest string
				mf.infra.EXPECT().JsonUnmarshal([]byte("abcd"), &dest).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_string",
			args: args{userID: 3},
			mockFields: func(mf mockFields) {
				mf.cache.EXPECT().Get(context.Background(), mockKey).Return("token", nil)
				var dest string
				mf.infra.EXPECT().JsonUnmarshal([]byte("token"), &dest).DoAndReturn(
					func(input []byte, dest interface{}) error {
						*dest.(*string) = "token"
						return nil
					})
			},
			want: "token",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				cache: NewMockredisRepoProvider(ctrl),
				infra: NewMockinfraRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			r := &Resource{
				cache: mockFields.cache,
				infra: mockFields.infra,
			}

			got, err := r.GetJWTFromCache(context.Background(), test.args.userID)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestResource_SetJWTToCache(t *testing.T) {
	mockKey := "account:jwt:3"
	mockConfig := &configuration.AppConfig{
		JWT: configuration.JWTConfig{
			TTL: 1,
		},
	}

	type mockFields struct {
		cache *MockredisRepoProvider
		infra *MockinfraRepoProvider
	}

	type args struct {
		userID int64
		jwt    string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_Set_error_then_return_error",
			args: args{
				userID: 3,
				jwt:    "token",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.cache.EXPECT().Set(context.Background(), mockKey, "token", time.Duration(1)*time.Second).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				userID: 3,
				jwt:    "token",
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().GetConfig().Return(mockConfig)
				mf.cache.EXPECT().Set(context.Background(), mockKey, "token", time.Duration(1)*time.Second).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				cache: NewMockredisRepoProvider(ctrl),
				infra: NewMockinfraRepoProvider(ctrl),
			}
			test.mockFields(mockFields)

			r := &Resource{
				cache: mockFields.cache,
				infra: mockFields.infra,
			}

			err := r.SetJWTToCache(context.Background(), test.args.userID, test.args.jwt)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
