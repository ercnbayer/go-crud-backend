package main

import (
	_ "go-backend/api"
	_ "go-backend/config"
	_ "go-backend/db"
	//"go-backend/migration"
)

func main() {
	//migration.Up(db.Db)
	//migration.Down(db.Db)
}
