package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/dracuxan/GoMongo/models"
)

type UserController struct {
	Database *mongo.Database
}

func NewUserController(db *mongo.Database) *UserController {
	return &UserController{Database: db}
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}
	u := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = uc.Database.Collection("user").FindOne(ctx, bson.M{"_id": objeID}).Decode(&u)

	return c.Status(fiber.StatusOK).JSON(u)
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	var users []models.User
	collection := uc.Database.Collection("user") // Ensure collection name matches
	filter := bson.D{}

	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to fetch users"})
	}
	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		var result models.User
		if err := curr.Decode(&result); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "failed to decode user"})
		}
		users = append(users, result)
	}

	if err := curr.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cursor error"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	return nil
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	return nil
}
