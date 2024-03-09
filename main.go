package main

import (
	"backend/database"
	"backend/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoute(app)

	//app.Listen(":4000")

	// Start the Fiber app on port 4000
	if err := app.Listen(":4000"); err != nil {
		log.Fatal(err)
	}
}
