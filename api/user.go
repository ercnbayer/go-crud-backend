package api

import (
	"go-backend/db"
	"go-backend/logger"

	"github.com/gofiber/fiber/v2"
)

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	if err := c.BodyParser(id); err != nil { // check if err
		logger.Error("GET SINGLE USER error = ", err)
		return c.SendStatus(404)
	}

	person := &db.Person{ID: id} // person instance

	person, err := db.ReadPerson(person, db.Db)
	if err != nil { //check if err is not null

		logger.Error(err.Error())
		return c.JSON(err.Error())
	}

	return c.JSON(person) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	person := &db.Person{ID: id}

	person, err := db.InsertPerson(person, db.Db) // for delete api

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id again

	person := &db.Person{ID: id} // creating instance

	if err := c.BodyParser(person); err != nil { // check if err
		logger.Error("error = ", err)
		return c.SendStatus(404)
	}

	person, err := db.PatchUpdatePerson(person, db.Db)

	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(person)

}
func createUser(c *fiber.Ctx) error {

	person := new(db.Person)

	if err := c.BodyParser(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.SendStatus(404)
	}

	person, err := db.InsertPerson(person, db.Db)

	if err != nil {

		logger.Error(" false request err", err.Error())
		return c.JSON(err.Error())

	}

	return c.JSON(person)

}
func listUsers(c *fiber.Ctx) error {

	people, err := db.ReturnAllQueries(db.Db) //list api

	if err != nil {

		logger.Error(" false request err", err.Error())
		return c.JSON(err)
	}

	return c.JSON(people)
}

func Init() {

	app := fiber.New()

	userApi := app.Group("/user") // grouping rotues

	userApi.Post("/", createUser) // creating user

	userApi.Get(":id", getSingleUser) // get single user

	userApi.Get("/", listUsers) //list all users

	userApi.Patch(":id", updateUser) //update user

	userApi.Delete(":id", deleteUser) //delete user

	logger.Fatal(app.Listen(":3000"))
}
