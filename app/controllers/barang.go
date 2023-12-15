package controllers

import (
	"go-fiber/app/models"
	"go-fiber/databases/conn"

	"github.com/gofiber/fiber/v2"
)

func IndexBarang(ctx *fiber.Ctx) error {
	var barang []models.Barang

	 conn.DB.Find(&barang)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data found",
		"barang": barang,
	})
}

func StoreBarang(ctx *fiber.Ctx) error {
	var request models.Barang

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	
	conn.DB.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data created successfully",
		"barang": request,
	})
}

func ShowBarang(ctx *fiber.Ctx) error {
	var barang models.Barang
	id := ctx.Params("id")

	result := conn.DB.First(&barang, id)

	if result.RowsAffected != 1 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data found",
		"barang": barang,
	})
}

func DeleteBarang(ctx *fiber.Ctx) error {
	var barang models.Barang
	id := ctx.Params("id")

	result := conn.DB.Where("id = ?", id).Delete(&barang)

	if result.RowsAffected != 1 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Berhasil dihapus",
	})
}

func UpdateBarang(c *fiber.Ctx) error {
	var barang models.Barang

	id := c.Params("id")

	if err := c.BodyParser(&barang); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	result := conn.DB.Where("id = ?", id)

	result.Updates(&barang)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found!",
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil di update",
		"barang": barang,
	})
}