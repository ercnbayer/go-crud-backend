package migration

//Migrations and Migrations_Arr structs are necessary for this file. Maybe you can define sorting method in here. But there is no other method needed in this file. You can move into utils other fonctions and structs
//Also you can manege tour up down logics in seperate files for better readablity
import (
	"go-backend/db"
	"go-backend/logger"
	"reflect"
	"sort"
)

// This definition not requeired. You can use like that inside migrations struct: Up   func() error
type Up func() error
type Down func() error

// Not plural. Should be singular naming
type Migrations struct {
	Name   string
	UpFn   Up
	DownFn Down
}
type CommittedMigration struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"name"`
}

func InsertMigration(Name string) error {

	if err := db.Db.Save(&CommittedMigration{Name: Name}).Error; err != nil {
		return err
	}

	logger.Info(Name, "has been saved")
	return nil
}

func DeleteMigration(Name string) error {

	//You are deleting the migration but you are deleting that from only migration table!!
	// You are not calling the down method of the migarion
	err := db.Db.Where("Name=?", Name).Delete(&CommittedMigration{}).Error

	if err != nil {
		logger.Info("delete err", Name)
		return err
	}

	logger.Info("deleted", Name)
	return nil
}

func (table *CommittedMigration) TableName() string {

	return "migrations"
}

var Migrations_Arr []Migrations

var Committed_Migs []CommittedMigration

func GetMigsFromDB() error {
	if err := db.Db.Find(&Committed_Migs).Error; err != nil {
		//check if err
		return err
	}

	logger.Info("commited migs get success")

	return nil
}

func SortMigArr() {

	sort.Slice(Migrations_Arr, func(i, j int) bool {

		v1 := reflect.ValueOf(Migrations_Arr[i]).FieldByName("Name") // sort according to name
		v2 := reflect.ValueOf(Migrations_Arr[j]).FieldByName("Name")
		return v1.String() < v2.String()
	})
}
func RunUp() {

	SortMigArr()

	//you can just call migElement as migration
	for _, migElement := range Migrations_Arr {

		//There is no error control. This causes when migarion not successfull state not prevents migration table insertion. And this causes you cant run the migration after fix that. Before at the table unsuccefull migartion saved as succefull
		migElement.UpFn()
		InsertMigration(migElement.Name) // insert

	}

}

func RunDown() error {
	//We are using this query but we dont used its response. Why you didnt used this querys result before deleting migrations
	//Also this code only deleting last executd migraition.
	//You need to have 2 ron down menthod
	if err := GetMigsFromDB(); err != nil {
		logger.Error("getCommittedMigErr")
		return err
	}

	if err := DeleteMigration(Committed_Migs[len(Committed_Migs)-1].Name); err != nil {

		return err
	}

	return nil

}

func init() {

	if !db.Db.Migrator().HasTable(&CommittedMigration{}) {
		if err := db.Db.Migrator().CreateTable(&CommittedMigration{}); err != nil {
			panic("failed to create table")
		}
	}

}

//mig2_Arr oluştur
// up ve down //*ptr func //alfabetik sıralı olacak //mig run burada olacak
