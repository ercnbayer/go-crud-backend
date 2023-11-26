package main

import (
	"go-backend/api"
	"go-backend/config"
	"go-backend/db"
)

func main() {

	err := config.Init(3) //LOG LEVEL

	if err == nil {

		err = db.Init()

		if err == nil {

			api.Init()

		}
	}

}
