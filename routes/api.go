package routes

import (
	"go-fiber/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(ctx *fiber.App) {
	route := ctx.Group("/api/v1")

	route.Post("/auth/register", controllers.Register)
}