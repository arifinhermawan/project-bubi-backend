package configuration

// AppConfig holds configuration needed for bubi
type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

// DatabaseConfig holds configuration related with database
type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
}
