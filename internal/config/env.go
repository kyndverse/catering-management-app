package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	Addr      string
	DBDSN     string
	LOG_LEVEL string
}

func NewEnvConfig() *ConfigEnv {
	_ = godotenv.Load()

	cfg := &ConfigEnv{
		Addr:      getEnv("ADDR", ":8080"),
		DBDSN:     os.Getenv("DBDSN"),
		LOG_LEVEL: os.Getenv("LOG_LEVEL"),
	}

	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
