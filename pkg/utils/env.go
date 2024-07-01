package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error occurred loading env variables")
	}
}
