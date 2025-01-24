package routes

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dracuxan/GoMongo/controllers"
)

var url string = "mongodb://127.0.0.1:27017/"

func getSession() *mongo.Client {
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error creating mongo client:", err)
	}

	return client
}

func Routes(app *fiber.App) {
	client := getSession()
	db := client.Database("GoMongo")
	uc := controllers.NewUserController(db)

	app.Get("/users", uc.GetUsers)
	app.Get("/user/:id", uc.GetUser)
	app.Post("/user", uc.CreateUser)
	app.Post("/user/:id", uc.UpdateUser)
	app.Delete("/user/:id", uc.DeleteUser)
}
