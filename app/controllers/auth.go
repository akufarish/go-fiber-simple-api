package controllers

import (
	"go-fiber/app/request"

	"github.com/gofiber/fiber/v2"
)

func Login(nama string) string {
	return "Hello " + nama
}

func Register(ctx *fiber.Ctx) error {
		var request request.RegisterRequest
		if err := ctx.BodyParser(&request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad Request",
			})
		}
		// Respond with success if everything went well
		return ctx.JSON(fiber.Map{
			"message": "Registration successful",
			"email": request.Email,
			"username": request.Username,
		})
}