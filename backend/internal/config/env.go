package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvIfLocal() {
	//command: export GO_ENV=development
	env := os.Getenv("GO_ENV")
	if env == "development" || env == "" {
		//command: export ENV_PATH=/Users/ave/kvizo/.env
		envPath := os.Getenv("ENV_PATH")
		var err error

		if envPath != "" {
			err = godotenv.Load(envPath)
			if err != nil {
				log.Printf("Failed to load .env from %s: %v", envPath, err)
			} else {
				log.Printf(".env loaded from %s for local development", envPath)
			}
		} else {
			err = godotenv.Load()
			if err != nil {
				log.Println("No .env file found in current directory, continuing with environment variables")
			} else {
				log.Println(".env loaded from current directory for local development")
			}
		}
	}
}
