package db

import (
	"go-backend/config"
	"go-backend/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() error {

	var err error

	Db, err = gorm.Open(postgres.Open(config.ConfigString), &gorm.Config{}) //connecting gorm

	if err != nil { //check if err
		logger.Fatal(err)
	}
	return nil
}
