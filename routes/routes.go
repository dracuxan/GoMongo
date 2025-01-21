package routes

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"

	"github.com/dracuxan/GoMongo/controllers"
)

var url string = "mongodb://localhost:27107"

func getSession() *mgo.Session {
	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	return s
}

func Routes(app *fiber.App) {
	uc := controllers.NewUserController(getSession())

	app.Get("/users", uc.GetUsers)
	app.Get("/user/:id", uc.GetUser)
	app.Post("/user", uc.CreateUser)
	app.Post("/user/:id", uc.UpdateUser)
	app.Delete("/user/:id", uc.DeleteUser)
}
