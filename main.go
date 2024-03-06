package main

import (
	"backend/database"
	"backend/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	router.SetupRoute(app)

	app.Listen(":4000")
}
