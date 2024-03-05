package auth

import (
	"backend/models"
	"backend/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Birthday struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type SignupInfo struct {
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  *Birthday `json:"birthday"`
	Address   string    `json:"address"`
	Username  string    `json:"username"`
	//Password  string    `json:"password"`
}

func Signup(c *fiber.Ctx) error {

	signupInfo := new(SignupInfo)

	if err := c.BodyParser(signupInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
			"data":    signupInfo,
			"error":   err,
		})
	}

	//users := utils.Mi.Db.Collection("users")
	userColl := utils.Mi.Db.Collection("users")
	user := new(models.User)

	//Check duplicate email
	if err := userColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupInfo.Email}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error while verifying user email address",
				"data":    signupInfo,
				"error":   err,
			})
		}
	}

	if user.Email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists",
			"data":    signupInfo,
		})
	}

	//Check duplicate phone

	//Check duplicate username

	//parse string time
	Birthtime := time.Date(signupInfo.Birthday.Year, time.Month(signupInfo.Birthday.Month), signupInfo.Birthday.Day, 0, 0, 0, 0, time.UTC)

	//Check age (if needed)

	//Save data

	doc := bson.D{{Key: "email", Value: signupInfo.Email}, {Key: "phone", Value: signupInfo.Phone}, {Key: "firstname", Value: signupInfo.Firstname}, {Key: "lastname", Value: signupInfo.Lastname}, {Key: "username", Value: signupInfo.Username}, {Key: "address", Value: signupInfo.Address}, {Key: "birthday", Value: Birthtime}}
	insertedUser, err := userColl.InsertOne(context.TODO(), doc)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while creating user",
			"data":    signupInfo,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Signup",
		"userID":  insertedUser.InsertedID,
	})
}
