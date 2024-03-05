package main

import (
	"backend/auth"
	"backend/utils"
	"log"

	"github.com/gofiber/fiber"
)

func main() {

	if err := utils.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	authRoute := app.Group("/auth")

	authRoute.Post("/signup", auth.Signup)

	app.Listen(":4000")
}
