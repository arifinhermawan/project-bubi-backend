package server

import (
	// internal package
	"io"
	"net/http"
	"time"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

//go:generate mockgen -source=infra.go -destination=infra_mock.go -package=server

type authenticationProvider interface {
	JWTAuthorization(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc
}

// configProvider provides methods available in config infra.
type configProvider interface {
	// GetConfig will get configuration that had been saved to memory.
	GetConfig() *configuration.AppConfig
}

// golangProvider provides methods available in golang infra.
type golangProvider interface {
	// GetTimeGMT7 will get current time in GMT+7
	GetTimeGMT7() time.Time

	// JsonMarshal returns the JSON encoding of input.
	JsonMarshal(input interface{}) ([]byte, error)

	// JsonUnmarshal parses the JSON-encoded data and stores the result in the value pointed to by dest.
	JsonUnmarshal(input []byte, dest interface{}) error
}

// readerProvider provides methods available in reader infra.
type readerProvider interface {
	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF. Because ReadAll is
	// defined to read from src until EOF, it does not treat an EOF from Read
	// as an error to be reported.
	ReadAll(input io.Reader) ([]byte, error)
}

// InfraParam represents parameters needed to initialize infrastructure.
type InfraParam struct {
	Auth   authenticationProvider
	Config configProvider
	Golang golangProvider
	Reader readerProvider
}

// Infra holds methods needed to initialize infrastructure.
type Infra struct {
	Auth   authenticationProvider
	Config configProvider
	Golang golangProvider
	Reader readerProvider
}

// NewInfra will initialize a new instance of Infra.
func NewInfra(param InfraParam) *Infra {
	return &Infra{
		Auth:   param.Auth,
		Config: param.Config,
		Golang: param.Golang,
		Reader: param.Reader,
	}
}

// GetConfig will get configuration that had been saved to memory.
func (infra *Infra) GetConfig() *configuration.AppConfig {
	return infra.Config.GetConfig()
}

// GetTimeGMT7 will get current time in GMT+7
func (infra *Infra) GetTimeGMT7() time.Time {
	return infra.Golang.GetTimeGMT7()
}

// JsonMarshal returns the JSON encoding of input.
func (infra *Infra) JsonMarshal(input interface{}) ([]byte, error) {
	return infra.Golang.JsonMarshal(input)
}

// JsonUnmarshal parses the JSON-encoded data and stores the result in the value pointed to by dest.
func (infra *Infra) JsonUnmarshal(input []byte, dest interface{}) error {
	return infra.Golang.JsonUnmarshal(input, dest)
}

func (infra *Infra) JWTAuthorization(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return infra.Auth.JWTAuthorization(endpointHandler)
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (infra *Infra) ReadAll(input io.Reader) ([]byte, error) {
	return infra.Reader.ReadAll(input)
}
