package utils

import (
	// golang package
	"testing"

	// external package
	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

func TestInitRedisConn(t *testing.T) {
	redisNewClientOri := redis.NewClient
	miniRedis, _ := miniredis.Run()
	mockClient := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})

	defer func() {
		redisNewClient = redisNewClientOri
		miniRedis.Close()
		mockClient.Close()
	}()

	type args struct {
		cfg *configuration.RedisConfig
	}
	tests := []struct {
		name       string
		args       args
		mock       func(opt *redis.Options) *redis.Client
		mockResult func() (string, error)
		want       *redis.Client
		wantErr    error
	}{
		{
			name: "when_failed_to_initialize_connection_then_return_error",
			args: args{
				cfg: &configuration.RedisConfig{
					Address:  "address",
					Password: "password",
				},
			},
			mock: func(opt *redis.Options) *redis.Client {
				client := redis.NewClient(&redis.Options{
					Addr: "address",
				})

				return client
			},
			wantErr: errInitRedisConn,
		},
		{
			name: "when_successfully_initialize_connection_then_return_nil_error",
			args: args{
				cfg: &configuration.RedisConfig{
					Address:  "address",
					Password: "password",
				},
			},
			mock: func(opt *redis.Options) *redis.Client {
				return mockClient
			},
			want: mockClient,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			redisNewClient = test.mock

			got, err := InitRedisConn(test.args.cfg)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
