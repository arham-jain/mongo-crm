package conf

type DatabaseConfig struct {
	ConnectionString string `json:"connectionString"`
	DatabaseName     string `json:"databaseName"`
}
