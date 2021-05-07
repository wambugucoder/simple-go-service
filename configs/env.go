package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ExtractEnvKey (Your env key)/*
func ExtractEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No Key Found for:" + key)
	}
	return os.Getenv(key)
}
