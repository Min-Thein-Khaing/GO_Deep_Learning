package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	DB_SSLMODE         string
	DATABASE_URL       string
	JWT_SECRET         string
	JWT_REFRESH_SECRET string
	PORT               string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		DB_HOST:            strings.TrimSpace(os.Getenv("DB_HOST")),
		DB_PORT:            strings.TrimSpace(os.Getenv("DB_PORT")),
		DB_USER:            strings.TrimSpace(os.Getenv("DB_USER")),
		DB_PASSWORD:        strings.TrimSpace(os.Getenv("DB_PASSWORD")),
		DB_NAME:            strings.TrimSpace(os.Getenv("DB_NAME")),
		DB_SSLMODE:         strings.TrimSpace(os.Getenv("DB_SSLMODE")),
		DATABASE_URL:       strings.TrimSpace(os.Getenv("DATABASE_URL")),
		JWT_SECRET:         strings.TrimSpace(os.Getenv("JWT_SECRET")),
		JWT_REFRESH_SECRET: strings.TrimSpace(os.Getenv("JWT_REFRESH_SECRET")),
		PORT:               strings.TrimSpace(os.Getenv("PORT")),
	}

	if cfg.DB_HOST == "" {
		cfg.DB_HOST = "localhost"
	}
	if cfg.DB_PORT == "" {
		cfg.DB_PORT = "5432"
	}
	if cfg.DB_SSLMODE == "" {
		cfg.DB_SSLMODE = "disable"
	}
	if cfg.PORT == "" {
		cfg.PORT = "8080"
	}

	if cfg.DATABASE_URL == "" && (cfg.DB_USER == "" || cfg.DB_NAME == "") {
		return Config{}, fmt.Errorf("missing required database configuration (DB_USER/DB_NAME or DATABASE_URL)")
	}

	if cfg.JWT_SECRET == "" || cfg.JWT_REFRESH_SECRET == "" {
		return Config{}, fmt.Errorf("missing required environment variables: JWT_SECRET or JWT_REFRESH_SECRET")
	}

	return cfg, nil
}

func (c *Config) PostgresDSN() string {
	if c.DATABASE_URL != "" {
		return c.DATABASE_URL
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_SSLMODE)
}

