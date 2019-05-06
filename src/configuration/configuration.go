package configuration

type (
	// Config
	Config struct {
		*DatabaseConfig
		*ServiceConfig
	}
)

// NewConfig
func NewConfig(databaseConfig *DatabaseConfig, serviceConfig *ServiceConfig) *Config {
	return &Config{ServiceConfig: serviceConfig, DatabaseConfig: databaseConfig}
}
