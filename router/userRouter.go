package router

import (
	"depen/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, handler *handler.UserHandler) {
	user := app.Group("/user")
	user.Post("/register", handler.RegisterUser)
	user.Post("/login", handler.LoginUser)
}
