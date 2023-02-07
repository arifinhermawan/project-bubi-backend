package utils

import (
	// golang package
	"context"
	"errors"
	"log"

	// external package
	"github.com/redis/go-redis/v9"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

var (
	errInitRedisConn = errors.New("failed to init redis connection!")

	redisNewClient = redis.NewClient
)

// InitDBConn will initialize connection to redis.
func InitRedisConn(cfg *configuration.RedisConfig) (*redis.Client, error) {
	client := redisNewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("[InitRedisConn] client.Ping().Result() got error: %+v\n", errInitRedisConn)
		return nil, errInitRedisConn
	}
	log.Println(pong)
	log.Println("successfully connect to redis")

	return client, nil
}
