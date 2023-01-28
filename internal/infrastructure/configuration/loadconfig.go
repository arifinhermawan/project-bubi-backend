package configuration

import (
	// golang package
	"errors"
	"fmt"
	"log"
	"os"

	// external package
	"github.com/spf13/viper"
)

var (
	configFile = "bubi.%s.yaml"

	// variables for error
	errConfigNotFound          = errors.New("config not found!")
	errConfigFailedToUnmarshal = errors.New("failed to unmarshal config!")

	// mock viper
	viperReadInConfig = viper.ReadInConfig
	viperUnmarshal    = viper.Unmarshal
)

// GetConfig will get configuration that had been saved to memory.
func (c Configuration) GetConfig() AppConfig {
	c.doLoadConfigOnce.Do(func() {
		cfg, err := c.LoadConfig()
		if err != nil {
			log.Fatalf("[GetConfig] c.LoadConfig() got an error: %+v\n", err)
		}

		c.Config = cfg
	})

	return c.Config
}

// LoadConfig will read configuration from config file.
func (c Configuration) LoadConfig() (AppConfig, error) {
	viper.AddConfigPath("files/")

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	meta := map[string]interface{}{
		"env": env,
	}

	configFile := fmt.Sprintf(configFile, env)
	meta["config_file"] = configFile

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err := viperReadInConfig()
	if err != nil {
		log.Printf("[LoadConfig] viperReadInConfig() got an error: %+v\nMeta: %+v", err, meta)
		return AppConfig{}, errConfigNotFound
	}

	var config AppConfig
	err = viperUnmarshal(&config)
	if err != nil {
		log.Printf("[LoadConfig] viperUnmarshal() got an error: %+v\nMeta: %+v", err, meta)
		return AppConfig{}, errConfigFailedToUnmarshal
	}

	return config, nil
}
