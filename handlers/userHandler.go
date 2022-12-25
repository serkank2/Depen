package handlers

import "github.com/gofiber/fiber/v2"

func LoginHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Login")
}

func RegisterHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Register")
}

func UpdateHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Update")

}
func DeleteHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete")

}
