package db

import (
	"errors"
	"go-backend/logger"
)

type Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email ;not null"`
}

// TableName overrides the table name used by User to `profiles`
func (table *Person) TableName() string {
	return "persons"
}

func InsertPerson(person *Person) error {

	if err := Db.Save(&person).Error; err != nil {
		return err
	}

	logger.Info(person)

	return nil

}

func DeletePerson(id string) (string, error) {

	var QueryResult = Db.Delete(&Person{ID: id})

	if err := QueryResult.Error; err != nil {

		logger.Info("delete ", err)
		return id, err

	}
	if QueryResult.RowsAffected == 0 {

		logger.Info("USER IS NOT FOUND")

		return id, errors.New("user NOT FOUND")
	}

	return id, nil
}

func ReadPerson(id string, person *Person) error {

	/*if err := Db.First(&Person{ID: id}).Error; err != nil {
		return err
	}*/

	person.ID = id

	if err := Db.First(person).Error; err != nil {
		return err
	}

	return nil
}

func PatchUpdatePerson(person *Person) error {

	var Result = Db.Updates(person)
	if err := Result.Error; err != nil {

		logger.Error("err update", err)
		return err

	}
	//Result = Result.Save(person)
	if err := Result.Error; err != nil {

		logger.Error("ERR UPDATE:", err)
		return err
	}

	if Result.RowsAffected == 0 {
		logger.Error("err user not found")

		return errors.New("user NOT FOUND")
	}

	// Return the updated person
	return nil

}

func GetUsers() ([]Person, error) {
	var people []Person // creating person arr
	if err := Db.Find(&people).Error; err != nil {
		//check if err
		return nil, err
	}

	return people, nil
}
