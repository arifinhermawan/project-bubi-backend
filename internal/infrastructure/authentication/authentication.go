package authentication

import (
	// golang package
	"encoding/json"
	"log"
	"net/http"
	"strings"

	// external package
	"github.com/golang-jwt/jwt"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

// configProvider holds all methods served by package configuration that will
// be needed by package authentication
type configProvider interface {
	// GetConfig will get configuration that had been saved to memory.
	GetConfig() *configuration.AppConfig
}

type Auth struct {
	cfg configProvider
}

// NewAuth will instantiate a new instance of Auth
func NewAuth(cfg *configuration.Configuration) *Auth {
	return &Auth{
		cfg: cfg,
	}
}

// JWTAuthorization will check authorization of a JWT.
func (auth *Auth) JWTAuthorization(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header["Authorization"]
		if len(authHeader) == 0 {
			log.Printf("[JWTAuthorization] authorization empty\n")
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("unauthorized!")
			return
		}

		sliced := strings.Split(authHeader[0], " ")
		token, err := jwt.Parse(sliced[1], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				log.Printf("[JWTAuthorization] failed to parse token method\n")
				writer.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(writer).Encode("unauthorized!")
				return nil, nil
			}
			return []byte(auth.cfg.GetConfig().JWT.Secret), nil
		})
		if err != nil {
			log.Printf("[JWTAuthorization] failed to parse jwt:%+v\n", err)
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("unauthorized!")
			return
		}

		if token.Valid {
			endpointHandler(writer, request)
		}
	})
}
