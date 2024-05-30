package initializers

import (
	"log"
	"server/database"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file: ", err)
	}
}

// LoadDB
func LoadDb() {
	database.InitDb()
}
