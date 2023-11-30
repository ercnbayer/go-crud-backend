package api

import (
	db "go-backend/db"
	"go-backend/logger"
	"go-backend/validator"

	"github.com/gofiber/fiber/v2"
)

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	var person = new(db.Person)
	person.ID = id
	_, err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.JSON(c.SendStatus(400), "invalid req")
	}

	person, err = db.ReadPerson(id)
	if err != nil { //check if err is not null

		logger.Error(err.Error(), err)
		return c.JSON(c.SendStatus(404), err.Error())
	}

	return c.JSON(person) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	_, err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.JSON(c.SendStatus(400), "invalid req")

	}

	id, err = db.DeletePerson(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.JSON(c.SendStatus(404), "USER NOT FOUND")
	}

	return c.JSON(id)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id again

	_, err := validator.ValidateID(id)

	if err != nil {
		return c.JSON(c.SendStatus(400), err.Error())
	}
	person := new(db.Person) // creating instance

	if err := c.BodyParser(person); err != nil { // check if err

		logger.Error("UPDATE USER error = ", err, person)

		return c.JSON(c.SendStatus(400), err.Error())
	}
	person.ID = id
	person, err = db.PatchUpdatePerson(person)

	if err != nil {
		return c.JSON(c.SendStatus(404), "USER NOT FOUND")
	}

	return c.JSON(person)

}
func createUser(c *fiber.Ctx) error {

	person := &db.Person{}

	if err := c.BodyParser(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.JSON(c.SendStatus(400), "false req")
	}

	person, err := db.InsertPerson(person)

	if err != nil {

		logger.Error(" false request err", err.Error())
		return c.JSON(c.SendStatus(400), "db err")

	}

	return c.JSON(person)

}
func listUsers(c *fiber.Ctx) error {

	people, err := db.GetUsers() //list api

	if err != nil {

		logger.Error(" no user found", err.Error())
		return c.JSON(c.SendStatus(404), "no user found")
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
		logger.Fatal("err")
	}
}
