package api

import (
	db "go-backend/db"
	"go-backend/logger"
	"go-backend/validator"

	"github.com/gofiber/fiber/v2"
)

type PersonPayload struct {
	ID       string
	Name     string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required"`
}

func mapPersonPayloadToDbPerson(person *PersonPayload) *db.Person {

	return &db.Person{ID: person.ID, Name: person.Name, Email: person.Email, Password: person.Password}
}

func mapPersonPayloadToDbPersonCreate(person *PersonPayload) *db.Person {

	return &db.Person{Name: person.Name, Email: person.Email, Password: person.Password}
}

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.JSON(c.SendStatus(400), err.Error())
	}

	person, err := db.ReadPerson(id)

	if err != nil { //check if err is not null

		logger.Error(err.Error(), err)
		return c.JSON(c.SendStatus(404), err.Error())
	}

	return c.JSON(person) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.JSON(c.SendStatus(400), err.Error())

	}

	_, err = db.DeletePerson(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.JSON(c.SendStatus(404), err.Error())
	}

	return c.JSON(id)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id again

	err := validator.ValidateID(id)

	if err != nil {
		return c.JSON(c.SendStatus(400), err.Error())
	}

	person := new(PersonPayload) // creating instance

	if err := c.BodyParser(person); err != nil { // check if err

		logger.Error("UPDATE USER error = ", err, person)

		return c.JSON(c.SendStatus(400), err.Error())
	}

	if err := validator.ValidateUpdatedStruct(person); err != nil {

		return c.JSON(c.SendStatus(400), err.Error())
	}

	person.ID = id

	dbPerson, err := db.PatchUpdatePerson(mapPersonPayloadToDbPerson(person))

	if err != nil {
		return c.JSON(c.SendStatus(404), err.Error())
	}

	return c.JSON(dbPerson)

}
func createUser(c *fiber.Ctx) error {

	person := new(PersonPayload)

	if err := c.BodyParser(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.JSON(c.SendStatus(400), err.Error())
	}

	if err := validator.ValidateStruct(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.JSON(c.SendStatus(400), err.Error())
	}

	dbPerson, err := db.InsertPerson(mapPersonPayloadToDbPersonCreate(person))

	if err != nil {

		logger.Error(" false request err", err.Error())
		return c.JSON(c.SendStatus(404), "user not founderr")

	}

	return c.JSON(dbPerson)

}
func listUsers(c *fiber.Ctx) error {

	people, err := db.GetUsers() //list api

	if err != nil {

		logger.Error(" no user found", err.Error())
		return c.JSON(c.SendStatus(404), err.Error())
	}

	return c.JSON(people)
}

func init() {

	app := fiber.New()

	userApi := app.Group("/user") // grouping rotues

	userApi.Post("/", createUser) // creating user

	userApi.Get(":id", getSingleUser) // get single user

	userApi.Get("/", listUsers) //list all users

	userApi.Patch(":id", updateUser) //update user

	userApi.Delete(":id", deleteUser) //delete user

	err := app.Listen(":3000")

	if err != nil {
		logger.Fatal("err", err)
	}
}
