package redis

import (
	// golang package
	"context"
	"testing"
	"time"

	// external package
	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepository_Del(t *testing.T) {
	type mockFields struct {
		redis redismock.ClientMock
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_Del_error_then_return_error",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectDel("keys").SetErr(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectDel("keys").SetVal(1)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			redis, mock := redismock.NewClientMock()
			mockFields := mockFields{
				redis: mock,
			}

			test.mockFields(mockFields)

			r := &RedisRepository{
				redis: redis,
			}

			err := r.Del(context.Background(), test.args.key)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestRedisRepository_Get(t *testing.T) {
	type mockFields struct {
		redis redismock.ClientMock
	}

	type args struct {
		key string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		want       string
		wantErr    error
	}{
		{
			name: "when_Exists_error_then_return_error",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectExists("keys").SetErr(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_Exists_return_0_then_return_empty_string",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectExists("keys").SetVal(0)
			},
		},
		{
			name: "when_Get_error_then_return_error",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectExists("keys").SetVal(1)
				mf.redis.ExpectGet("keys").SetErr(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil_error",
			args: args{key: "keys"},
			mockFields: func(mf mockFields) {
				mf.redis.ExpectExists("keys").SetVal(1)
				mf.redis.ExpectGet("keys").SetVal("abc")
			},
			want: "abc",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			redis, mock := redismock.NewClientMock()
			mockFields := mockFields{
				redis: mock,
			}
			test.mockFields(mockFields)

			r := &RedisRepository{
				redis: redis,
			}

			got, err := r.Get(context.Background(), test.args.key)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestRedisRepository_Set(t *testing.T) {
	type mockFields struct {
		redis redismock.ClientMock
		infra *MockinfraProvider
	}

	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_JsonMarshal_error_then_return_error",
			args: args{value: "abcd"},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().JsonMarshal("abcd").Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_Set_error_then_return_error",
			args: args{
				key:        "keys",
				value:      "abcd",
				expiration: 1,
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().JsonMarshal("abcd").Return([]byte("abcd"), nil)
				mf.redis.ExpectSet("keys", []byte("abcd"), 1).SetErr(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{
				key:        "keys",
				value:      "abcd",
				expiration: 1,
			},
			mockFields: func(mf mockFields) {
				mf.infra.EXPECT().JsonMarshal("abcd").Return([]byte("abcd"), nil)
				mf.redis.ExpectSet("keys", []byte("abcd"), 1).SetVal("")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			redis, mock := redismock.NewClientMock()

			mockFields := mockFields{
				infra: NewMockinfraProvider(ctrl),
				redis: mock,
			}
			test.mockFields(mockFields)

			r := &RedisRepository{
				redis: redis,
				infra: mockFields.infra,
			}

			err := r.Set(context.Background(), test.args.key, test.args.value, test.args.expiration)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
