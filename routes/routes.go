package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dracuxan/GoMongo/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.Home)
}
