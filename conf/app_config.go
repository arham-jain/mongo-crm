package conf

// AppConfig to unmarshal the config file
type AppConfig struct {
	DatabaseConfig DatabaseConfig `json:"databaseConfig"`
}

var appConfig = &AppConfig{}

// RegisterConfig global instance of config
func RegisterConfig(config *AppConfig) {
	appConfig = config
}

// Config to access the global app config
func Config() *AppConfig {
	return appConfig
}
