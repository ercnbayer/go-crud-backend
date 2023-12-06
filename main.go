package main

import (
	//_ "go-backend/api"
	_ "go-backend/config"
	_ "go-backend/db"
	"go-backend/migration"
	//"go-backend/migration"
)

func main() {

	migration.RunDown()
}
