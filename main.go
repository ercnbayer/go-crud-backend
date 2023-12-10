package main

import (
	//_ "go-backend/api"
	_ "go-backend/db"
	_ "go-backend/migration"
	migrationutils "go-backend/migration-utils"
	//"go-backend/migration"
)

func main() {
	//migration.RunUp()
	//migration.RunDownMigration("Person20231209003351")
	migrationutils.RunUpMigration("Person20231209002731")
}
