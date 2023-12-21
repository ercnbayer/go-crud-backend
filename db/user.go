package db

import (
	"errors"
	"go-backend/logger"
)

type Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	//You dont need to define column there. Gorm automaticly converting Name to name
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

//Rename this function to DeletePersonById
func DeletePerson(id string) (string, error) {

	var QueryResult = Db.Delete(&Person{ID: id})

	if err := QueryResult.Error; err != nil {

		//Make this more meaningfull message. Something went wrong when deleting Person by Id; err
		//Also this should be logger.Error, not logger.Info
		logger.Info("delete ", err)
		return id, err

	}

	if QueryResult.RowsAffected == 0 {

		//This should be warning at the same time make your all log messages standart look. Dont make an message ALL CHARACTERS CAPITAL, or all of them lower case. Log messages needs to be easy to read
		logger.Info("USER IS NOT FOUND")

		// Move your standart errors to your custom error package. And use this predefined errors
		return id, errors.New("user NOT FOUND")
	}

	return id, nil
}

//Why did you give id as extra perameter. You can just add this parameter 
//Rename this function to FindPersonById
func ReadPerson(id string, person *Person) error {

	//Remove unecessary comments. You can lookup your old code from commit history. This is uneccessary and meaningless for other contributers
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

	// Why did you named this as Result instead of result
 	var Result = Db.Updates(person)
	if err := Result.Error; err != nil {

		logger.Error("err update", err)
		return err

	}
	//Result = Result.Save(person)
	// Remove error control. You are not running another db query this is unecessary
	if err := Result.Error; err != nil {

		logger.Error("ERR UPDATE:", err)
		return err
	}

	if Result.RowsAffected == 0 {
		logger.Error("err user not found")
		//Use this error from common place to.
		return errors.New("user NOT FOUND")
	}

	// Return the updated person
	return nil

}

func GetUsers() ([]Person, error) {
	//Get this Person array from outside of the function by pointer. This should be standart in project if we gonna return sturct
	var people []Person // creating person arr
	if err := Db.Find(&people).Error; err != nil {
		//Unecessary commentline. You dont need to explain this :D
		//check if err
		return nil, err
	}

	return people, nil
}
