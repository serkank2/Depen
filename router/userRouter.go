package router

import "github.com/gofiber/fiber/v2"

func UserRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, User!")
	})
}
