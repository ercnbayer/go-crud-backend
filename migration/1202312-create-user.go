package migration

import (
	"gorm.io/gorm"
)

type Person struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"name;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email ;not null"`
}

func (table *Person) TableName() string {
	return "persons"
}

func Up(db *gorm.DB) error {
	// Use the db.Migrator().CreateTable method to create a new table
	err := db.Migrator().CreateTable(&Person{})
	if err != nil {
		return err
	}

	// Add columns or make other schema changes as needed

	return nil
}

func Down(db *gorm.DB) error {
	// Use the db.Migrator().DropTable method to drop the previously created table
	err := db.Migrator().DropTable(&Person{})
	if err != nil {
		return err
	}
	return nil
}
