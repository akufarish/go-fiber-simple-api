package handlers

import (
	"go-fiber/app/request"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(ctx *fiber.Ctx) error {
	// Parse the request body or parameters if needed
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
	})
}
