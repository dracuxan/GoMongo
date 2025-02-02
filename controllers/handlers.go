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

	err = uc.Database.Collection("users").FindOne(ctx, bson.M{"_id": objeID}).Decode(&u)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	var users []models.User
	collection := uc.Database.Collection("users") // Ensure collection name matches
	filter := bson.M{}                            // Empty filter to fetch all documents

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ensure you're properly calling Find with the correct context
	curr, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to fetch users", "details": err.Error()})
	}

	// Defer closing the cursor using the same context instead of background
	defer curr.Close(ctx) // Changed from context.Background() to ctx

	// Iterate over the cursor using the same context
	for curr.Next(ctx) { // Changed to use the correct ctx here
		var result models.User
		if err := curr.Decode(&result); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "failed to decode user"})
		}
		users = append(users, result)
	}

	// Check for any errors after iterating over the cursor
	if err := curr.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "cursor error"})
	}

	// Return the list of users as a JSON response
	return c.Status(fiber.StatusOK).JSON(users)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "unable to parse json body"})
	}

	if user.Id.IsZero() {
		user.Id = primitive.NewObjectID()
	}

	coll := uc.Database.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.InsertOne(ctx, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to insert user", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": res.InsertedID})
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	fid := c.Params("id")

	objId, err := primitive.ObjectIDFromHex(fid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	updateUser := make(map[string]interface{})

	if err = c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Unable to parse json body"})
	}

	if len(updateUser) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "No fields provided for update"})
	}
	coll := uc.Database.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": updateUser}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to update user", "details": err.Error()})
	}

	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "User updated successfully",
		"matchedCount":  res.MatchedCount,
		"modifiedCount": res.ModifiedCount,
	})
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	coll := uc.Database.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objId}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to delete user"})
	}

	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"message": "User deleted successfully", "deleteCount": res.DeletedCount})
}
