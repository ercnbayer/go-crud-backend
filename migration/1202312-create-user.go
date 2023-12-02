package migration

import (
	"go-backend/db"
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

func init() {

	db.Db.Set("gorm::table_options", "ENGINE=InnoDB").Migrator().CreateTable(Person{})
}
