package main

import (
	"depen/auth"
	"depen/config"
	"depen/handler"
	"depen/repo"
	"depen/router"
	"depen/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	accessToken, _ := auth.CreateToken("serkankaplan42@yandex.com")
	claims, _ := auth.ParseToken(accessToken)
	fmt.Println(claims["email"])
	// ----------------- User -----------------

	UserCollection := config.GetCollection(config.DB, "user") // create collection
	UserRepo := repo.NewUserRepo(UserCollection)              // create repo
	UserServices := services.NewUserServices(UserRepo)        // create services
	UserHandler := handler.NewUserHandler(UserServices)       // create handler
	router.UserRouter(app, UserHandler)                       // create router

	// ----------------- User -----------------

	app.Listen(":8080")
}
