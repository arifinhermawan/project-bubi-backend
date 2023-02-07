package redis

import (
	// golang package
	"context"
	"log"
	"time"
)

// Del will delete a key in redis.
func (repo *RedisRepository) Del(ctx context.Context, key string) error {
	redisInt := repo.redis.Del(ctx, key)
	_, err := redisInt.Result()
	if err != nil {
		meta := map[string]interface{}{
			"key": key,
		}

		log.Printf("[Del] redisInt.Result() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}

// Get will get the value of a redis key.
// First it will check whether the key exist or not.
// If it exists, then it will return the value.
// Otherwise, it returns empty string.
func (repo *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	meta := map[string]interface{}{
		"key": key,
	}

	redisInt := repo.redis.Exists(ctx, key)
	isExist, err := redisInt.Result()
	if err != nil {
		log.Printf("[Get] redisInt.Result() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	if isExist == 0 {
		return "", nil
	}

	redisString := repo.redis.Get(ctx, key)
	result, err := redisString.Result()
	if err != nil {
		log.Printf("[Get] redisString.Result() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	return result, nil
}

// Set will save the value of a key to redis.
func (repo *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	meta := map[string]interface{}{
		"key":        key,
		"expiration": expiration,
	}

	bytes, err := repo.infra.JsonMarshal(value)
	if err != nil {
		log.Printf("[Set] repo.infra.JsonMarshal() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	redisStatus := repo.redis.Set(ctx, key, bytes, expiration)
	_, err = redisStatus.Result()
	if err != nil {
		log.Printf("[Set] redisStatus.Result() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}
