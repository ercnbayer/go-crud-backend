package db

import (
	"go-backend/logger"
)

type Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"name;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email ;not null"`
}

// TableName overrides the table name used by User to `profiles`
func (table *Person) TableName() string {
	return "persons"
}

func InsertPerson(person *Person) (*Person, error) {

	if err := Db.Save(&person).Error; err != nil {
		return nil, err
	}

	logger.Info(person)

	return person, nil

}

func DeletePerson(id string) (string, error) {

	if err := Db.First(&Person{ID: id}).Error; err != nil {
		return id, err
	}

	if err := Db.Delete(&Person{ID: id}).Error; err != nil {

		return id, err

	}

	return id, nil
}

func ReadPerson(id string) (*Person, error) {

	if err := Db.First(&Person{ID: id}).Error; err != nil {
		return nil, err
	}

	var person = new(Person)
	person.ID = id

	if err := Db.Find(person).Error; err != nil {
		return nil, err
	}

	return person, nil
}

func PatchUpdatePerson(person *Person) (*Person, error) {

	existingPerson := &Person{ID: person.ID}

	if err := Db.Find(existingPerson).Error; err != nil { // if id does not exist
		return nil, err
	}

	if err := Db.Model(existingPerson).Updates(person).Error; err != nil {

		logger.Error("err update")
		return nil, err

	}

	// Return the updated person
	return existingPerson, nil

}

func GetUsers() ([]Person, error) {
	var people []Person // creating person arr
	if err := Db.Find(&people).Error; err != nil {
		//check if err
		return nil, err
	}

	return people, nil
}
