package migration

import (
	"go-backend/db"
	"go-backend/logger"
)

type Person20231209003351 struct {
	Age int
}

func (table Person20231209003351) TableName() string {
	return "persons"
}
func PersonUp20231209003351() error {

	if err := db.Db.Migrator().AddColumn(&Person20231209003351{}, "Age"); err != nil {
		logger.Error("Person20231209003351 ADD COLUMN ERR")
		return err
	}

	return nil
}
func PersonDown20231209003351() error {

	if err := db.Db.Migrator().DropColumn(&Person20231209003351{}, "Age"); err != nil {
		logger.Error(" Person20231209003351 Drop COLUMN ERR")
		return err
	}

	/*if err := InsertMigration(time.Now().String()); err != nil {
		return err
	}*/

	logger.Info("DROP COLUMN SUCCESS")

	return nil
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "Person20231209003351",
		UpFn:   PersonUp20231209003351,
		DownFn: PersonDown20231209003351,
	})
	logger.Info("TABLE INIT")

}
