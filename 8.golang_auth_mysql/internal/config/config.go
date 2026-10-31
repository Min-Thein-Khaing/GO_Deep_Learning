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
	MYSQL_DSN          string
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
		MYSQL_DSN:          strings.TrimSpace(os.Getenv("MYSQL_DSN")),
		JWT_SECRET:         strings.TrimSpace(os.Getenv("JWT_SECRET")),
		JWT_REFRESH_SECRET: strings.TrimSpace(os.Getenv("JWT_REFRESH_SECRET")),
		PORT:               strings.TrimSpace(os.Getenv("PORT")),
	}

	if cfg.DB_HOST == "" {
		cfg.DB_HOST = "localhost"
	}
	if cfg.DB_PORT == "" {
		cfg.DB_PORT = "3306"
	}
	if cfg.DB_USER == "" {
		cfg.DB_USER = "root"
	}
	if cfg.DB_NAME == "" {
		cfg.DB_NAME = "go_auth_mysql"
	}
	if cfg.PORT == "" {
		cfg.PORT = "8080"
	}

	if cfg.MYSQL_DSN == "" && (cfg.DB_USER == "" || cfg.DB_NAME == "") {
		return Config{}, fmt.Errorf("missing required database configuration (DB_USER/DB_NAME or MYSQL_DSN)")
	}

	if cfg.JWT_SECRET == "" || cfg.JWT_REFRESH_SECRET == "" {
		return Config{}, fmt.Errorf("missing required environment variables: JWT_SECRET or JWT_REFRESH_SECRET")
	}

	return cfg, nil
}

func (c *Config) MySQLDSN() string {
	if c.MYSQL_DSN != "" {
		return c.MYSQL_DSN
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME)
}

