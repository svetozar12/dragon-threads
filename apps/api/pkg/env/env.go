package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                       string
	POSTGRES_CONNECTION_STRING string
	GITHUB_CLIENT_ID           string
	GITHUB_SECRET              string
	GITHUB_REDIRECT_URL        string
	FRONTEND_URL               string
}

var Envs Config

/*
InitConfig initializes the configuration by parsing environment variables and storing them in Config and ServicesConfig structs.
*/
func InitConfig() {
	// Load environment variables from .env file
	godotenv.Load("apps/api/.env")

	Envs = Config{
		Port:                       getEnv("PORT", "3333"),
		POSTGRES_CONNECTION_STRING: getEnv("POSTGRES_CONNECTION_STRING", "postgres://postgres:postgrespw@172.17.0.1:5432"),
		GITHUB_CLIENT_ID:           getEnv("GITHUB_CLIENT_ID", ""),
		GITHUB_SECRET:              getEnv("GITHUB_SECRET", ""),
		GITHUB_REDIRECT_URL:        getEnv("GITHUB_REDIRECT_URL", "http://localhost:3333/v1/auth/callback"),
		FRONTEND_URL:               getEnv("FRONTEND_URL", "http://localhost:4200"),
	}

	fmt.Println("Envs were successfully loaded!")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
