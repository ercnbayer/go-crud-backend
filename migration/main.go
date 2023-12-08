package migration

import (
	"go-backend/logger"
	migrationutils "go-backend/migration-utils"
	"reflect"
	"sort"
)

type Up func() error
type Down func() error

type Migrations struct {
	Name   string
	UpFn   Up
	DownFn Down
}

var Migrations_Arr []Migrations

func SortMigArr() {

	sort.Slice(Migrations_Arr, func(i, j int) bool {

		v1 := reflect.ValueOf(Migrations_Arr[i]).FieldByName("Name") // sort according to name
		v2 := reflect.ValueOf(Migrations_Arr[j]).FieldByName("Name")
		return v1.String() < v2.String()
	})
}

func RunUp() {

	SortMigArr()

	for _, migElement := range Migrations_Arr {
		if err := migrationutils.SearchMigration(migElement.Name); err != nil {
			migElement.UpFn()
			migrationutils.InsertMigration(migElement.Name) // insert
			logger.Info("INSERTED NEW MIGRATION", migElement.Name)
		}

	}

}

func RunDown() error {

	if err := migrationutils.GetMigsFromDB(); err != nil {
		logger.Error("getCommittedMigErr")
		return err
	}

	if err := migrationutils.DeleteMigration(migrationutils.CommittedMigs[len(migrationutils.CommittedMigs)-1].Name); err != nil {

		return err
	}

	return nil

}

func init() {

}

//mig2_Arr oluştur
// up ve down //*ptr func //alfabetik sıralı olacak //mig run burada olacak
