package main

import (
	"depen/config"
	"depen/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Env()
	app := fiber.New()

	routers.Router(app)

	app.Listen(":" + config.GetPort())
}
