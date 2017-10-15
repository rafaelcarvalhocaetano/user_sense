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
			DatabaseURL: getEnv("DATABASE_URL", "postgres://root:postgres@localhost:5432/ex_typesense?sslmode=disable"),
		},
		Typesense: TypesenseConfig{
			Host:   getEnv("TYPESENSE_HOST", "http://localhost:8108"),
			APIKey: getEnv("TYPESENSE_API_KEY", "xyz"),
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
