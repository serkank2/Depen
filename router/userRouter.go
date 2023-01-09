package router

import (
	"depen/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/register", handler.RegisterUser)
	user.Post("/login", handler.RegisterUser)
}
