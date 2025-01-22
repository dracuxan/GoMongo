package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dracuxan/GoMongo/models"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if !bson.IsObjectIdHex(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user id not found"})
	}
	fid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.Session.DB("GoMongo").C("users").FindId(fid).One(&u); err != nil {
		c.Status(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	return nil
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
