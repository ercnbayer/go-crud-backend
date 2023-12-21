package api

import (
	db "go-backend/db"
	"go-backend/logger"
	"go-backend/validator"

	"github.com/gofiber/fiber/v2"
)

// Ypu need to create two seperate payloads for create and update. Before while creation user all fields and required except id but for update they are optional
type PersonPayload struct {
	ID       string
	Name     string `validate:"omitempty,required"`
	Password string `validate:"omitempty,required"`
	Email    string `validate:"omitempty,required,email"`
}

func mapPersonPayloadToDbPerson(person *PersonPayload, dbPerson *db.Person) {

	dbPerson.ID = person.ID
	dbPerson.Name = person.Name
	dbPerson.Email = person.Email
	dbPerson.Password = person.Password
}

func mapPersonPayloadToDbPersonCreate(person *PersonPayload, dbPerson *db.Person) {

	dbPerson.Name = person.Name
	dbPerson.Email = person.Email
	dbPerson.Password = person.Password

}

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)

		return c.Status(400).JSON(err.Error())
	}

	var person db.Person
	err = db.ReadPerson(id, &person)

	if err != nil { //check if err is not null

		logger.Error(err.Error(), err)

		return c.Status(404).JSON(err.Error())
	}

	//Is the status code 200. İf not add it
	return c.JSON(person) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	_, err = db.DeletePerson(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(id)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id again

	err := validator.ValidateID(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	//You can create body parse and validate as one function ->
	var person PersonPayload // creating instance

	if err := c.BodyParser(&person); err != nil { // check if err

		logger.Error("UPDATE USER error = ", err, person)

		return c.Status(404).JSON(err.Error())
	}

	person.ID = id
	var dbPerson db.Person

	if err := validator.ValidateUpdatedStruct(person); err != nil {

		logger.Error("Validate USER error = ", err, person)

		return c.Status(400).JSON(err.Error())
	}
	//-> until there

	mapPersonPayloadToDbPerson(&person, &dbPerson)
	err = db.PatchUpdatePerson(&dbPerson)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(dbPerson)

}
func createUser(c *fiber.Ctx) error {

	var person PersonPayload

	var dbPerson db.Person

	if err := c.BodyParser(&person); err != nil { //check if err//Unecessary :D
		logger.Error(" false request err", err)
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.ValidateStruct(&person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.Status(400).JSON(err.Error())
	}

	mapPersonPayloadToDbPersonCreate(&person, &dbPerson)

	err := db.InsertPerson(&dbPerson)

	if err != nil {

		logger.Error(" false request err", err.Error())
		return c.Status(404).JSON(err.Error())

	}

	return c.JSON(dbPerson)

}
func listUsers(c *fiber.Ctx) error {

	people, err := db.GetUsers() //list api

	if err != nil {

		logger.Error(" no user found", err.Error())
		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(people)
}

func UserInit() {

	userApi := App.Group("/user") // grouping rotues

	userApi.Post("/", createUser) // creating user

	userApi.Get(":id", getSingleUser) // get single user

	userApi.Get("/", listUsers) //list all users

	userApi.Patch(":id", updateUser) //update user

	userApi.Delete(":id", deleteUser) //delete user

}
