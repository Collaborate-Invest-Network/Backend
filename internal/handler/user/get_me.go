package user

import (
	"backend/database"
	"backend/internal/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *fiber.Ctx) error {
	userID, _ := primitive.ObjectIDFromHex(c.Params("userId"))

	usersColl := database.Mi.Db.Collection(database.UserCollection)
	user := new(models.User)
	filter := bson.D{{Key: "_id", Value: userID}}

	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "user not found",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get user",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func GetMe(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	usersColl := database.Mi.Db.Collection(database.UserCollection)
	user := new(models.User)
	filter := bson.D{{Key: "_id", Value: userId}}

	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "user not found",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "get current user",
		"data": fiber.Map{
			"user": user,
		},
	})

}
