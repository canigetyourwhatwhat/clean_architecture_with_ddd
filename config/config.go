package config

type (
	Config struct {
		Port string   `yaml:"port"`
		DB   Database `yaml:"database"`
	}

	Database struct {
		DatabaseName string `yaml:"databaseName"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
	}
)
