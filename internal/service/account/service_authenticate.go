package account

import (
	// external package
	"context"
	"log"
	"time"

	// external package
	"github.com/golang-jwt/jwt"
)

var (
	jwtNewWithClaims = jwt.NewWithClaims
)

// GenerateJWT will generate a new JWT for user if not exist.
// If it already exist, it will get the value from cache.
// If it's not exist, it will generate new token and save it to cache.
func (svc *Service) GenerateJWT(ctx context.Context, userID int64, email string) (string, error) {
	meta := map[string]interface{}{
		"user_id": userID,
		"email":   email,
	}

	existingJWT, err := svc.rsc.GetJWTFromCache(ctx, userID)
	if err != nil {
		log.Printf("[GenerateJWT] svc.rsc.GetJWTFromCache() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	if existingJWT != "" {
		return existingJWT, nil
	}

	expired := svc.infra.GetConfig().Account.ExpiredTimeInHour
	token := jwtNewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": svc.infra.GetTimeGMT7().Add(time.Hour * time.Duration(expired)).Unix(),
	})

	sec := []byte(svc.infra.GetConfig().JWT.Secret)
	tokenString, err := token.SignedString(sec)
	if err != nil {
		log.Printf("[GenerateJWT] token.SignedString() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	err = svc.rsc.SetJWTToCache(ctx, userID, tokenString)
	if err != nil {
		log.Printf("[GenerateJWT] svc.rsc.SetJWTToCache() got an error: %+v\nMeta:%+v\n", err, meta)
	}

	return tokenString, nil
}

// InvalidateJWT will delete user's cached JWT.
func (svc *Service) InvalidateJWT(ctx context.Context, userID int64) error {
	err := svc.rsc.DeleteJWTInCache(ctx, userID)
	if err != nil {
		meta := map[string]interface{}{
			"user_id": userID,
		}

		log.Printf("[InvalidateJWT] svc.rsc.DeleteJWTInCache() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}
