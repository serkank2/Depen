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
	c.BodyParser(&model)
	result := u.UserServices.RegisterUser(&model)
	return c.Status(200).JSON(result)
}
func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Login User",
	})
}

func NewUserHandler(userServices *services.UserServices) *UserHandler {
	return &UserHandler{UserServices: userServices}
}
