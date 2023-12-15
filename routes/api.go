package routes

import (
	"go-fiber/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(ctx *fiber.App) {	
	route := ctx.Group("/api/v1")

	route.Post("/auth/register", controllers.Register)
	route.Post("/auth/login", controllers.Login)

	route.Get("/barang", controllers.IndexBarang)
	route.Post("/barang", controllers.StoreBarang)
	route.Get("/barang/:id", controllers.ShowBarang)
	route.Delete("/barang/:id", controllers.DeleteBarang)
	route.Put("/barang/:id", controllers.UpdateBarang)
}