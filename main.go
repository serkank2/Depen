package main

import (
	router "depen/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.UserRouter(app)
	app.Listen(":8080")
}
