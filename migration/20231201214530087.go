package migration

import (
	"go-backend/db"
	"go-backend/logger"
)

type D1202312Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"name;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email ;not null"`
}

var entryName string

func (table D1202312Person) TableName() string {
	return "persons"
}
func D1202312PersonUp() error {
	// Use the db.Migrator().CreateTable method to create a new table
	if err := db.Db.Migrator().CreateTable(&D1202312Person{}); err != nil {
		logger.Info("Table init err")
		return err
	}
	logger.Info("table init  SUCCESS")
	/*if err := InsertMigration(entryName); err != nil {
		return err
	}*/

	// Add columns or make other schema changes as needed
	return nil
}
func D1202312PersonDown() error {
	// Use the db.Migrator().DropTable method to drop the previously created table
	if err := db.Db.Migrator().DropTable(&D1202312Person{}); err != nil {
		logger.Info("err drop table")
		return err
	}

	/*if err := DeleteMigration(); err != nil {
		return err
	}*/

	logger.Info("drop table success")

	return nil
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migrations{
		Name:   "20231201214530087",
		UpFn:   D1202312PersonUp,
		DownFn: D1202312PersonDown,
	})
	logger.Info("TABLE INIT")

}
