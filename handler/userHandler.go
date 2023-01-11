package handler

import (
	"depen/model"
	"depen/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserServices *services.UserServices
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var model model.RegisterUser
	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Body Parse Error",
		})
	}
	result := u.UserServices.RegisterUser(&model)
	return c.Status(result.HttpStatusCode).JSON(result)
}
func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	var model model.LoginUser
	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Body Parse Error",
		})
	}
	result := u.UserServices.LoginUser(&model)
	return c.Status(result.HttpStatusCode).JSON(result)
}

func NewUserHandler(userServices *services.UserServices) *UserHandler {
	return &UserHandler{UserServices: userServices}
}
