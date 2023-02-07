package redis

import (
	// golang package
	"context"
	"time"

	// external package
	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=./repo.go -destination=./repo_mock.go -package=redis

// infraProvider holds all methods from infra that will be used in package redis.
type infraProvider interface {
	// JsonMarshal returns the JSON encoding of input.
	JsonMarshal(input interface{}) ([]byte, error)
}

// redisProvider holds all methods from redis that will be used in package redis.
type redisProvider interface {
	// Del will delete a key in redis.
	Del(ctx context.Context, keys ...string) *redis.IntCmd

	// Exists will check whether a key is exist in redis.
	Exists(ctx context.Context, keys ...string) *redis.IntCmd

	// Get will get the value of a redis key.
	Get(ctx context.Context, key string) *redis.StringCmd

	// Set will save the value of a key to redis.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

// RedisRepositoryParam holds all parameters needed to instansiate new
// RedisRepository instance
type RedisRepositoryParam struct {
	Infra infraProvider
	Redis redisProvider
}

type RedisRepository struct {
	infra infraProvider
	redis redisProvider
}

// NewRedisRepository will create a new instance of RedisRepository.
func NewRedisRepository(param RedisRepositoryParam) *RedisRepository {
	return &RedisRepository{
		infra: param.Infra,
		redis: param.Redis,
	}
}
