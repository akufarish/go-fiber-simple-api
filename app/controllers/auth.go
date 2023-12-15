package controllers

import (
	"go-fiber/app/models"
	"go-fiber/databases/conn"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error)  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Login(ctx *fiber.Ctx) error {
	var request models.User

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	var user models.User
	result := conn.DB.Where("email = ? OR username = ?", request.Email, request.Username).First(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Password Salah",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}


func Register(ctx *fiber.Ctx) error {
    var request models.User
	var user models.User

    if err := ctx.BodyParser(&request); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Bad Request",
        })
    }

	if conn.DB.Where("email = ?", request.Email).First(&user) != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email sudah terdaftar!",
		})
	}

    hashedPassword, err := Hash(request.Password)

    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal server error",
        })
    }

	request.Password = hashedPassword

    conn.DB.Create(&request)

    return ctx.JSON(fiber.Map{
        "message": "Registration successful",
    })
}
