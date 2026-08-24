package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGODB_URI   string
	MONGO_DB_NAME string
	JWT_SECRET    string
}

func Load() (Config, error) {

	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := Config{
		MONGODB_URI:   strings.TrimSpace(os.Getenv("MONGODB_URI")),
		MONGO_DB_NAME: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWT_SECRET:    strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}
	if cfg.MONGODB_URI == "" || cfg.MONGO_DB_NAME == "" || cfg.JWT_SECRET == "" {
		return Config{}, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
