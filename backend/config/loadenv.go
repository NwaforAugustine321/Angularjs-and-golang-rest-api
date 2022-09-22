package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Can not load env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
