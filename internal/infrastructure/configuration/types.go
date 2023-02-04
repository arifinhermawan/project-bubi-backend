package configuration

// AppConfig holds configuration needed for bubi
type AppConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

// DatabaseConfig holds configuration related with database
type DatabaseConfig struct {
	Driver         string `mapstructure:"driver"`
	Host           string `mapstructure:"host"`
	Name           string `mapstructure:"name"`
	Password       string `mapstructure:"password"`
	Port           int    `mapstructure:"port"`
	User           string `mapstructure:"user"`
	DefaultTimeout int    `mapstructure:"default_timeout_in_seconds"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}
