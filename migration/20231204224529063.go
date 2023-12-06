package migration

import (
	"go-backend/db"
	"go-backend/logger"
)

type D4202312Person struct {
	Age int
}

func (table D4202312Person) TableName() string {
	return "persons"
}

func D4202312PersonUp() error {
	// Use the db.Migrator().CreateTable method to create a new table
	if err := db.Db.Migrator().AddColumn(&D4202312Person{}, "Age"); err != nil {
		logger.Error("D402312 ADD COLUMN ERR")
		return err
	}

	/*if err := InsertMigration(time.Now().String()); err != nil {
		return err
	}*/
	logger.Info(" column ADD SUCCESS.")
	// Add columns or make other schema changes as needed
	return nil
}
func D4202312PersonDown() error {
	// Use the db.Migrator().CreateTable method to create a new table
	if err := db.Db.Migrator().DropColumn(&D4202312Person{}, "Age"); err != nil {
		logger.Error("D402312 Drop COLUMN ERR")
		return err
	}

	/*if err := InsertMigration(time.Now().String()); err != nil {
		return err
	}*/

	logger.Info("DROP COLUMN SUCCESS")
	// Add columns or make other schema changes as needed
	return nil
}
func init() {
	Migrations_Arr = append(Migrations_Arr, Migrations{
		Name:   "20231204224529063",
		UpFn:   D4202312PersonUp,
		DownFn: D4202312PersonDown,
	})
	logger.Info("alter init")

}
