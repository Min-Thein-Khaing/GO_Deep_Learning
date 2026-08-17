package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGODB_URI   string
	MONGO_DB_NAME string
	ServerPORT    string
}

func Load() (Config, error) {

	//godotenv.Load() this call in project have .env testing if work including .env not have not working
	
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load .env file: %v", err)
	}
	mongodbURI, err := extractEnv("MONGODB_URI")
	if err != nil {
		return Config{}, err
	}
	mongoDBName, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}
	serverPort, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}
	return Config{MONGODB_URI: mongodbURI, MONGO_DB_NAME: mongoDBName, ServerPORT: serverPort}, nil
}

func extractEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}
