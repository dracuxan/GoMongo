package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"

	"github.com/dracuxan/GoMongo/controllers"
)

var url string = "mongodb://127.0.0.1:27017/"

func getSession() *mgo.Session {
	s, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
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
