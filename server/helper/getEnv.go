package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env files")
	}
	log.Println(os.Getenv(key))
	return os.Getenv(key)
}
