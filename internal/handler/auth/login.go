package auth

import (
	"backend/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginBody struct {
	Username string `json:"username" validate: "required"`
	Password string `json:"password" validate: "required"`
}

func Login(c *fiber.Ctx) error {
	loginBody := new(LoginBody)

	if err := c.BodyParser(loginBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"data":    nil,
			"message": err.Error(),
		})
	}

	//validate input
	errors := utils.ValidateStruct(loginBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data": fiber.Map{
				"errors": errors,
			},
		})
	}

	user := new(models.User)
	usersColl := database.Mi.Db.Collection(database.UserCollection)

	//verify (email/phone/username)
	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: loginBody.Username}},
			bson.D{{Key: "phone", Value: loginBody.Username}},
			bson.D{{Key: "username", Value: loginBody.Username}},
		},
		},
	}
	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Incorrect credentials",
				"data":    nil,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "erroe",
				"message": err.Error(),
				"data":    nil,
			})
		}
	}

	//check password hash
	if match := utils.CheckPasswordHash(loginBody.Password, user.Password); !match {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Incorrect credentials",
			"data":    nil,
		})
	}

	//update user active status
	filterUser := bson.D{{Key: "_id", Value: user.Id}}
	updateUser := bson.D{{Key: "4set", Value: bson.D{{Key: "isActive", Value: true}}}}
	_, err := usersColl.UpdateOne(context.TODO(), filterUser, updateUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "erroe",
			"message": err.Error(),
			"data":    nil,
		})
	}

	//fetch updated user
	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	//JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("SECRECT"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "You have logged in successfully",
		"data": fiber.Map{
			"user":       user,
			"accesToken": t,
		},
	})
}
