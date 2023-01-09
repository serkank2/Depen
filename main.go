package main

import (
	"depen/config"
	"depen/handler"
	"depen/repo"
	router "depen/router"
	"depen/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// ----------------- User -----------------

	userCollection := config.GetCollection(config.DB, "user") // create collection
	UserRepo := repo.NewUserRepo(userCollection)              // create repo
	UserServices := services.NewUserServices(UserRepo)        // create services
	UserHandler := handler.NewUserHandler(UserServices)       // create handler
	router.UserRouter(app, UserHandler)                       // create router

	// ----------------- User -----------------

	app.Listen(":8080")
}
