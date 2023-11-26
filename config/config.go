package config

import (
	"fmt"
	"go-backend/logger"
	"os"

	"github.com/joho/godotenv"
)

var ConfigString string

func Init(logLvl uint8) error {

	logger.LogLevel = logLvl

	err := godotenv.Load() //load env
	if err != nil {
		logger.Fatal("Error loading .env file")
		return err
	}

	dbName := os.Getenv("POSTGRES_DB") // getting env vars
	userName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("PORT")

	ConfigString = fmt.Sprintf("host=localhost user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Etc/UTC", userName, dbPassword, dbName, dbPort)

	return nil
}
