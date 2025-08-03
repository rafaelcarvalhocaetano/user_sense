package config

import (
	"os"
)

type Config struct {
	Database  DatabaseConfig
	Typesense TypesenseConfig
	Server    ServerConfig
}

type DatabaseConfig struct {
	DatabaseURL string
}

type TypesenseConfig struct {
	Host   string
	APIKey string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			DatabaseURL: getEnv("DATABASE_URL", "postgres://root:postgres@postgres-sfs:5432/user-sense?sslmode=disable"),
		},
		Typesense: TypesenseConfig{
			Host:   getEnv("TYPESENSE_HOST", "http://typesense-sfs:8108"),
			APIKey: getEnv("TYPESENSE_API_KEY", "xpto"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func (c *Config) GetServerAddress() string {
	return ":" + c.Server.Port
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
