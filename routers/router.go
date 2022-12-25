package routers

import (
	"depen/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Get("/user/login", handlers.LoginHandler)
	app.Get("/user/register", handlers.RegisterHandler)
	app.Get("/user/update", handlers.UpdateHandler)
	app.Get("/user/Delete", handlers.DeleteHandler)

}
