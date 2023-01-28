package configuration

import "sync"

// Configuration holds "properties" owned by package configuration
type Configuration struct {
	Config AppConfig

	doLoadConfigOnce *sync.Once
}

// NewConfiguration initialize new instance of Configuration.
func NewConfiguration() *Configuration {
	return &Configuration{
		doLoadConfigOnce: new(sync.Once),
	}
}
