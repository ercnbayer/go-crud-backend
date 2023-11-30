package migration

import (
	"go-backend/db"
)

func init() {

	db.Db.Set("gorm::table_options", "ENGINE=InnoDB").Migrator().CreateTable(&db.Person{})
}
