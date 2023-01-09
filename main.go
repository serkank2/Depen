package main

import (
	"context"
	"depen/config"
	router "depen/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	userCollection := config.GetCollection(config.DB, "user")
	_, err := userCollection.InsertOne(context.Background(), bson.M{"name": "John", "age": 56})
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	router.UserRouter(app)
	app.Listen(":8080")
}
