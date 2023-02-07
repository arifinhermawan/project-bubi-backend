package redis

import (
	// golang package
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewRedisRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfra := NewMockinfraProvider(ctrl)
	mockRedis := NewMockredisProvider(ctrl)

	want := &RedisRepository{
		infra: mockInfra,
		redis: mockRedis,
	}
	assert.Equal(t, want, NewRedisRepository(RedisRepositoryParam{
		Infra: mockInfra,
		Redis: mockRedis,
	}))
}
