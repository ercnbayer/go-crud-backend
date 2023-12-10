package migration

import (
	"go-backend/db"
	"go-backend/logger"
)

type Person20231209002731 struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"name;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email ;not null"`
}

func (table Person20231209002731) TableName() string {
	return "persons"
}
func PersonUp20231209002731() error {
	// Use the db.Migrator().CreateTable method to create a new table
	if err := db.Db.Migrator().CreateTable(&Person20231209002731{}); err != nil {
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

func PersonDown20231209002731() error {

	if err := db.Db.Migrator().DropTable(&Person20231209002731{}); err != nil {
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

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "Person20231209002731",
		UpFn:   PersonUp20231209002731,
		DownFn: PersonDown20231209002731,
	})
	logger.Info("TABLE INIT")

}
