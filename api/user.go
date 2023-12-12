package api

import (
	db "go-backend/db"
	"go-backend/logger"
	"go-backend/validator"

	"github.com/gofiber/fiber/v2"
)

// You can create name this struct like PersonPayload or PersonDTO. Bu PersonPayload is more suitable naming
type Person struct {
	ID       string
	Name     string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required"`
}

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	//Add this to your validation struct like `validate:"required,uuid4"`
	err := validator.ValidateID(id)

	if err != nil {

		logger.Error("invalid req", err)
		//You should return error messages from validation
		return c.JSON(c.SendStatus(400), "invalid req")
	}

	var person = new(db.Person)
	person.ID = id

	person, err = db.ReadPerson(id)

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

	err := validator.ValidateID(id)

	if err != nil {
		return c.JSON(c.SendStatus(400), err.Error())
	}
	person := new(Person) // creating instance

	if err := c.BodyParser(person); err != nil { // check if err

		logger.Error("UPDATE USER error = ", err, person)

		return c.JSON(c.SendStatus(400), err.Error())
	}

	if err := validator.ValidateUpdatedStruct(person); err != nil {

		return c.JSON(c.SendStatus(400), err.Error())
	}

	person.ID = id

	//Dont map api.Person to db.Person like that. You can create seperate function for this purpose. You can creat this function like mapPersonPayloadToPerson(d *api.PersonPaylaod) db.Person { your convertation stuff in here} make this for all function in al layers. You can manage your struct convertions like this by mapping functions
	dbPerson, err := db.PatchUpdatePerson(&db.Person{ID: person.ID, Name: person.Name, Email: person.Email, Password: person.Password})

	if err != nil {
		return c.JSON(c.SendStatus(404), "USER NOT FOUND")
	}

	return c.JSON(dbPerson)

}
func createUser(c *fiber.Ctx) error {

	person := &Person{}

	if err := c.BodyParser(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.JSON(c.SendStatus(400), "false req")
	}

	if err := validator.ValidateStruct(person); err != nil { //check if err
		logger.Error(" false request err", err)
		return c.JSON(c.SendStatus(400), "false req")
	}

	dbPerson, err := db.InsertPerson(&db.Person{Name: person.Name, Email: person.Email, Password: person.Password})

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
