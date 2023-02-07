package configuration

// AppConfig holds configuration needed for bubi
type AppConfig struct {
	Account  AccountConfig  `mapstructure:"account"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

// ------------------------------
// | structs for infrastructure |
// ------------------------------

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

// DatabaseConfig holds configuration related with redis
type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

// ----------------------
// | structs for config |
// ----------------------

type AccountConfig struct {
	ExpiredTimeInHour int `mapstructure:"expired_times_in_hour"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	TTL    int    `mapstructure:"ttl_in_seconds"`
}
