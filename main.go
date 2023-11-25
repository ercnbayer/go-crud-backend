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

func UpdatePerson(person *Person, db *gorm.DB) {

	db.Save(person) // where condition istiyor id e göre hareket etmiyor

	fmt.Println("updated:", person)

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

	if err := db.First(person).Error; err != nil {
		return nil, err
	}

	if err := db.Model(person).Updates(person).Error; err != nil {
		return nil, err
	}

	// Return the updated person
	return person, nil

}

func ReturnAllQueries(db *gorm.DB) ([]Person, error) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {

		return nil, err
	}

	return people, nil
}

func listUsers(c *fiber.Ctx) error {

	people, err := ReturnAllQueries(db)

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(people)
}

func getSingleUser(c *fiber.Ctx) error {

	// Get the ID from the URL parameter
	id := c.Params("id")

	person := &Person{ID: id}

	person, err := ReadPerson(person, db)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id")

	person := &Person{ID: id}

	person, err := DeletePerson(person, db)

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	person := &Person{ID: id}

	if err := c.BodyParser(person); err != nil {
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

	if err := c.BodyParser(person); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	person, err := InsertPerson(person, db)

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func main() {

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("POSTGRES_DB")

	userName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Etc/UTC", userName, dbPassword, dbName, dbPort)

	fmt.Println(dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userApi := app.Group("/user")

	userApi.Post("/", createUser)

	// getting user if no error
	userApi.Get(":id", getSingleUser)

	userApi.Get("/", listUsers)

	userApi.Patch(":id", updateUser)

	userApi.Delete(":id", deleteUser)

	log.Fatal(app.Listen(":3000"))

}
