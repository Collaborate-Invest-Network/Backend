package auth

import (
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

type SignupInfo struct {
	email     string `json:"email"`
	phone     string `json:"phone"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	birthday  string `json:"birthday"`
	address   string `json:"address"`
	username  string `json:"username"`
	password  string `json:"password"`
}

func Signup(c *fiber.Ctx) error {

	signupInfo := new(SignupInfo)

	if err := c.BodyParser(signupInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
			"data":    signupInfo,
		})
	}

	//users := utils.Mi.Db.Collection("users")
	_ = utils.Mi.Db.Collection("users")

	return c.JSON(fiber.Map{
		"message": "Signup",
	})
}
