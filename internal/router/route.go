package router

import (
	"backend/internal/handler/auth"
	"backend/internal/handler/user"
	"backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	authRoute := app.Group("/auth")
	authRoute.Post("/signup", auth.Signup)
	authRoute.Post("/login", auth.Login)

	userRoute := app.Group("/user")
	userRoute.Get("/me", middlewares.Protected(), user.GetMe)
	userRoute.Get("/:userId", user.GetUser)
}
