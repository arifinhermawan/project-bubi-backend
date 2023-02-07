package account

import (
	// golang package
	"context"
	"log"
	"strconv"
	"time"
)

const (
	redisKeyJWT = "account:jwt:"
)

// DeleteJWTInCache will delete JWT of a user.
func (rsc *Resource) DeleteJWTInCache(ctx context.Context, userID int64) error {
	key := buildJWTKey(userID)

	meta := map[string]interface{}{
		"key": key,
	}

	err := rsc.cache.Del(ctx, key)
	if err != nil {
		log.Printf("[GetJWTFromCache] rsc.cache.Del() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}

// GetJWTFromCache will fetch JWT of a user from cache.
// If the ke doesn't exist, it will return empty string
func (rsc *Resource) GetJWTFromCache(ctx context.Context, userID int64) (string, error) {
	key := buildJWTKey(userID)

	meta := map[string]interface{}{
		"key": key,
	}

	redisJWT, err := rsc.cache.Get(ctx, key)
	if err != nil {
		log.Printf("[GetJWTFromCache] rsc.cache.Get() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	if redisJWT == "" {
		return "", nil
	}

	var jwt string
	err = rsc.infra.JsonUnmarshal([]byte(redisJWT), &jwt)
	if err != nil {
		log.Printf("[GetJWTFromCache] rsc.infra.JsonUnmarshal() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	return jwt, nil
}

// SetJWTToCache will save jwt of a user in cache.
func (rsc *Resource) SetJWTToCache(ctx context.Context, userID int64, jwt string) error {
	key := buildJWTKey(userID)
	ttl := rsc.infra.GetConfig().JWT.TTL

	meta := map[string]interface{}{
		"key": key,
		"ttl": ttl,
	}

	ttlDuration := time.Second * time.Duration(ttl)
	err := rsc.cache.Set(ctx, key, jwt, ttlDuration)
	if err != nil {
		log.Printf("[SetJWTToCache] rsc.cache.Set() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}

// buildJWTKey will build a redis key for jwt related case
func buildJWTKey(userID int64) string {
	userIDStr := strconv.FormatInt(userID, 10)
	return redisKeyJWT + userIDStr
}
