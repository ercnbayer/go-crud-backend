package main

import (
	//_ "go-backend/api"
	_ "go-backend/db"
	"go-backend/migration"
	_ "go-backend/migration"
	//"go-backend/migration"
)

func main() {
	migration.RunUp()
}
