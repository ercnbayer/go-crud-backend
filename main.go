package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

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

	fmt.Println(person, "inserted")

	return person, nil

}

func DeletePerson(person *Person, db *gorm.DB) (*Person, error) {

	if err := db.First(person).Error; err != nil {
		return nil, err
	}
	//db.Delete(&person) // where clause istiyor id ye göre hareket etmiyor
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}
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

	fmt.Println("found", person)

	return person, nil
}

func PatchUpdatePerson(person *Person, db *gorm.DB) (*Person, error) {

	if err := db.First(person).Error; err != nil { // if id does not exist
		return nil, err
	}

	if err := db.Model(person).Updates(person).Error; err != nil {
		return nil, err
	}

	// Return the updated person
	return person, nil

}

func ReturnAllQueries(db *gorm.DB) ([]Person, error) {
	var people []Person // creating person arr
	if err := db.Find(&people).Error; err != nil {
		//check if err
		return nil, err
	}

	return people, nil
}

func listUsers(c *fiber.Ctx) error {

	people, err := ReturnAllQueries(db) //list api

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(people)
}

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	person := &Person{ID: id} // person instance

	person, err := ReadPerson(person, db)
	if err != nil { //check if err is not null
		return c.JSON(err.Error())
	}

	return c.JSON(person) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	person := &Person{ID: id}

	person, err := DeletePerson(person, db) // for delete api

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id again

	person := &Person{ID: id} // creating instance

	if err := c.BodyParser(person); err != nil { // check if err
		fmt.Println("error = ", err)
		return c.SendStatus(404)
	}

	person, err := PatchUpdatePerson(person, db)

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}
func createUser(c *fiber.Ctx) error {

	person := new(Person)

	if err := c.BodyParser(person); err != nil { //check if err
		fmt.Println("error = ", err)
		return c.SendStatus(404)
	}

	person, err := InsertPerson(person, db)

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func main() {

	app := fiber.New() // fiber inst

	err := godotenv.Load() //load env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("POSTGRES_DB") // getting env vars
	userName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Etc/UTC", userName, dbPassword, dbName, dbPort)

	fmt.Println(dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //connecting gorm

	if err != nil { //check if err
		log.Fatal(err)
	}

	userApi := app.Group("/user") // grouping rotues

	userApi.Post("/", createUser) // creating user

	userApi.Get(":id", getSingleUser) // get single user

	userApi.Get("/", listUsers) //list all users

	userApi.Patch(":id", updateUser) //update user

	userApi.Delete(":id", deleteUser) //delete user

	log.Fatal(app.Listen(":3000"))

}
