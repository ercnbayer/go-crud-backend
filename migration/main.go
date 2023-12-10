package migration

//Migrations and Migrations_Arr structs are necessary for this file. Maybe you can define sorting method in here. But there is no other method needed in this file. You can move into utils other fonctions and structs. In this case you can run up and down methods
//Also you can manege tour up down logics in seperate files for better readablity
import (
	"go-backend/logger"
	migrationutils "go-backend/migration-utils"
	"reflect"
	"sort"
)

// This definitions not required. You can use like that inside migrations struct: UpFn   func() error
type Up func() error
type Down func() error

// Not plural. Should be singular naming
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

			//There is no error control. This causes when migarion not successfull state not prevents migration table insertion. And this causes you cant run the migration after fix that. Before at the table unsuccefull migartion saved as succefull
			migElement.UpFn()
			migrationutils.InsertMigration(migElement.Name) // insert
			logger.Info("INSERTED NEW MIGRATION", migElement.Name)
		}

	}

}

func RunDown() error {

	//We are using this query but we dont used its response. Why you didnt used this querys result before deleting migrations
	//Also this code only deleting last executd migraition.
	//You need to have 2 ron down menthod
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
