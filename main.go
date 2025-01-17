package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dracuxan/GoMongo/routes"
)

func main() {
	port := ":1313"
	app := fiber.New()
	routes.Routes(app)
	app.Listen(port)
}
