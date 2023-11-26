package db

import (
	"go-backend/logger"

	"gorm.io/gorm"
)

var Db *gorm.DB

type Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

// TableName overrides the table name used by User to `profiles`
func (table *Person) TableName() string {
	return "persons"
}

func InsertPerson(person *Person, db *gorm.DB) (*Person, error) {

	if err := db.Save(&person).Error; err != nil {
		return nil, err
	}

	logger.Info(person)

	return person, nil

}

func DeletePerson(person *Person, db *gorm.DB) (*Person, error) {

	if err := db.First(person).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&person).Error; err != nil {

		return nil, err

	}

	return person, nil
}

func ReadPerson(person *Person, db *gorm.DB) (*Person, error) {

	if err := db.First(person).Error; err != nil {
		return nil, err
	}

	if err := db.Find(person).Error; err != nil {
		return nil, err
	}

	return person, nil
}

func PatchUpdatePerson(person *Person, db *gorm.DB) (*Person, error) {

	existingPerson := &Person{ID: person.ID}

	if err := db.Find(existingPerson).Error; err != nil { // if id does not exist
		return nil, err
	}

	if err := db.Model(existingPerson).Updates(person).Error; err != nil {
		logger.Error("err update")
		return nil, err
	}

	// Return the updated person
	return existingPerson, nil

}

func ReturnAllQueries(db *gorm.DB) ([]Person, error) {
	var people []Person // creating person arr
	if err := db.Find(&people).Error; err != nil {
		//check if err
		return nil, err
	}

	return people, nil
}
